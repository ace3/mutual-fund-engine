package grpc

import (
	"context"
	"errors"

	"github.com/ace3/mutual-fund-engine/internal/models"
	"github.com/ace3/mutual-fund-engine/internal/utils"
	pb "github.com/ace3/mutual-fund-engine/pkg/pb/nobi"
	"gorm.io/gorm"
)

type NobiInvestmentHandler struct {
	pb.UnimplementedNobiInvestmentServiceServer
	db *gorm.DB
}

func NewNobiInvestmentHandler(db *gorm.DB) *NobiInvestmentHandler {
	return &NobiInvestmentHandler{db: db}
}

// Implement gRPC methods here, e.g.:
func (h *NobiInvestmentHandler) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	if req.GetName() == "" || req.GetUsername() == "" {
		return nil, errors.New("name and username are required")
	}
	// Check if username already exists
	var existing models.User
	err := h.db.Where("username = ?", req.GetUsername()).First(&existing).Error
	if err == nil {
		return nil, errors.New("username already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	user := models.User{
		Name:     req.GetName(),
		Username: req.GetUsername(),
	}
	if err := h.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &pb.AddUserResponse{UserId: int32(user.ID)}, nil
}

func (h *NobiInvestmentHandler) UpdateTotalBalance(ctx context.Context, req *pb.UpdateTotalBalanceRequest) (*pb.UpdateTotalBalanceResponse, error) {
	if req.GetCurrentBalance() < 0 {
		return nil, errors.New("current_balance must be >= 0")
	}
	var totalUnit float64
	h.db.Model(&models.UserInvestment{}).Select("SUM(unit)").Scan(&totalUnit)
	var nab float64
	if totalUnit == 0 {
		nab = 1.0
	} else {
		nab = req.GetCurrentBalance() / totalUnit
	}
	// Round down to 4 decimals
	nab = utils.RoundDown(nab, 4)
	// Save to nab_history
	h.db.Table("nab_history").Create(&models.NABHistory{NAB: nab})
	return &pb.UpdateTotalBalanceResponse{NabAmount: nab}, nil
}

func (h *NobiInvestmentHandler) ListNAB(ctx context.Context, req *pb.ListNABRequest) (*pb.ListNABResponse, error) {
	var nabs []models.NABHistory
	h.db.Table("nab_history").Order("created_at desc").Find(&nabs)
	resp := &pb.ListNABResponse{}
	for _, n := range nabs {
		resp.Nabs = append(resp.Nabs, &pb.NABEntry{
			Nab:  n.NAB,
			Date: n.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return resp, nil
}

func (h *NobiInvestmentHandler) TopUp(ctx context.Context, req *pb.TopUpRequest) (*pb.TopUpResponse, error) {
	if req.GetAmountRupiah() <= 0 {
		return nil, errors.New("amount_rupiah must be > 0")
	}
	var user models.User
	if err := h.db.First(&user, req.GetUserId()).Error; err != nil {
		return nil, errors.New("user not found")
	}
	// Get latest NAB
	var nab models.NABHistory
	if err := h.db.Table("nab_history").Order("created_at desc").First(&nab).Error; err != nil {
		nab.NAB = 1.0
	}
	unit := req.GetAmountRupiah() / nab.NAB
	unit = utils.RoundDown(unit, 4)
	var inv models.UserInvestment
	err := h.db.Where("user_id = ?", user.ID).First(&inv).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		inv = models.UserInvestment{UserID: user.ID, Unit: 0}
	}
	inv.Unit += unit
	if err == nil {
		h.db.Save(&inv)
	} else {
		h.db.Create(&inv)
	}
	// Record transaction
	h.db.Create(&models.Transaction{
		UserID: user.ID,
		Type:   "TOPUP",
		Amount: req.GetAmountRupiah(),
		Unit:   unit,
		NABAt:  nab.NAB,
	})
	// Calculate total balance
	totalBalance := utils.RoundDown(inv.Unit*nab.NAB, 2)
	return &pb.TopUpResponse{
		NilaiUnitHasilTopup: unit,
		NilaiUnitTotal:      inv.Unit,
		SaldoRupiahTotal:    totalBalance,
	}, nil
}

func (h *NobiInvestmentHandler) Withdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {
	if req.GetAmountRupiah() <= 0 {
		return nil, errors.New("amount_rupiah must be > 0")
	}
	var user models.User
	if err := h.db.First(&user, req.GetUserId()).Error; err != nil {
		return nil, errors.New("user not found")
	}
	// Get latest NAB
	var nab models.NABHistory
	if err := h.db.Table("nab_history").Order("created_at desc").First(&nab).Error; err != nil {
		nab.NAB = 1.0
	}
	var inv models.UserInvestment
	if err := h.db.Where("user_id = ?", user.ID).First(&inv).Error; err != nil {
		return nil, errors.New("user has no investment")
	}
	unitToWithdraw := req.GetAmountRupiah() / nab.NAB
	unitToWithdraw = utils.RoundDown(unitToWithdraw, 4)
	if unitToWithdraw > inv.Unit {
		return nil, errors.New("withdrawal exceeds owned unit")
	}
	inv.Unit -= unitToWithdraw
	h.db.Save(&inv)
	// Record transaction
	h.db.Create(&models.Transaction{
		UserID: user.ID,
		Type:   "WITHDRAW",
		Amount: req.GetAmountRupiah(),
		Unit:   unitToWithdraw,
		NABAt:  nab.NAB,
	})
	totalBalance := utils.RoundDown(inv.Unit*nab.NAB, 2)
	return &pb.WithdrawResponse{
		NilaiUnitSetelahWithdraw: unitToWithdraw,
		NilaiUnitTotal:           inv.Unit,
		SaldoRupiahTotal:         totalBalance,
	}, nil
}

func (h *NobiInvestmentHandler) ListMembers(ctx context.Context, req *pb.ListMembersRequest) (*pb.ListMembersResponse, error) {
	var nabs models.NABHistory
	h.db.Table("nab_history").Order("created_at desc").First(&nabs)
	currentNAB := nabs.NAB
	if currentNAB == 0 {
		currentNAB = 1.0
	}

	type MemberResult struct {
		UserID int
		Unit   float64
	}

	page := int32(0)
	limit := int32(20)
	if req.Page != nil {
		page = req.GetPage()
	}
	if req.Limit != nil {
		limit = req.GetLimit()
	}

	var results []MemberResult
	query := h.db.Table("users u").
		Select("u.id as user_id, COALESCE(ui.unit, 0) as unit").
		Joins("LEFT JOIN user_investments ui ON u.id = ui.user_id").
		Order("u.id asc").
		Offset(int(page) * int(limit)).
		Limit(int(limit))
	if req.UserId != nil {
		query = query.Where("u.id = ?", req.GetUserId())
	}
	query.Find(&results)

	resp := &pb.ListMembersResponse{}
	for _, m := range results {
		resp.Members = append(resp.Members, &pb.MemberEntry{
			UserId:                  int32(m.UserID),
			TotalUnitPerUid:         m.Unit,
			TotalAmountrupiahPerUid: utils.RoundDown(m.Unit*currentNAB, 2),
			CurrentNab:              currentNAB,
		})
	}
	return resp, nil
}
