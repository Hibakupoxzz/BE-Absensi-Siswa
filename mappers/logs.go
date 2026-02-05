package mappers

import (
	"github.com/KicauOrgspark/BE-Absensi-Siswa/dto/responses"
	"github.com/KicauOrgspark/BE-Absensi-Siswa/models"
)

func ToLogsResponse(l models.AttedanceLogs) responses.LogsRes {
	return responses.LogsRes{
		ID: l.ID,
		User: responses.UserMini{
			ID:       l.User.ID,
			FullName: l.User.FullName,
		},
		Status:      l.Status,
		CapturedIp:  l.CapturedIp,
		ClockInTime: l.ClockInTime,
	}
}

func ListToLogsResponse(l []models.AttedanceLogs) []responses.LogsRes {
	res := make([]responses.LogsRes, 0, len(l))
	for _, ls := range l {
		res = append(res, ToLogsResponse(ls))
	}

	return res
}
