package stats

import (
	"errors"
	"testing"

	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/model"

	"go.uber.org/mock/gomock"
)

func TestStatsService_GetStats(t *testing.T) {
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		repository func() adapters.StatsRepository
	}
	tests := []struct {
		name      string
		fields    fields
		wantInt1  int
		wantInt2  int
		wantLimit int
		wantStr1  string
		wantStr2  string
		wantHits  int
		wantErr   bool
	}{
		{
			name: "successful stats retrieval",
			fields: fields{
				repository: func() adapters.StatsRepository {
					m := adapters.NewMockStatsRepository(ctrl)
					m.EXPECT().GetMostFrequentRequest().Return(
						&model.StatsResult{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz", Hits: 10}, nil).
						Times(1)
					return m
				},
			},
			wantInt1:  3,
			wantInt2:  5,
			wantLimit: 15,
			wantStr1:  "Fizz",
			wantStr2:  "Buzz",
			wantHits:  10,
			wantErr:   false,
		},
		{
			name: "error in stats retrieval",
			fields: fields{
				repository: func() adapters.StatsRepository {
					m := adapters.NewMockStatsRepository(ctrl)
					m.EXPECT().GetMostFrequentRequest().Return(nil, errors.New("failed")).Times(1)
					return m
				},
			},
			wantInt1:  0,
			wantInt2:  0,
			wantLimit: 0,
			wantStr1:  "",
			wantStr2:  "",
			wantHits:  0,
			wantErr:   true,
		},
		{
			name: "repository returns zero values",
			fields: fields{
				repository: func() adapters.StatsRepository {
					m := adapters.NewMockStatsRepository(ctrl)
					m.EXPECT().GetMostFrequentRequest().Return(nil, nil).Times(1)
					return m
				},
			},
			wantInt1:  0,
			wantInt2:  0,
			wantLimit: 0,
			wantStr1:  "",
			wantStr2:  "",
			wantHits:  0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStats(
				tt.fields.repository(),
			)

			sts, err := s.GetStats()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if sts.Int1 != tt.wantInt1 {
					t.Errorf("GetStats() Int1 = %v, want %v", sts.Int1, tt.wantInt1)
				}
				if sts.Int2 != tt.wantInt2 {
					t.Errorf("GetStats() Int2 = %v, want %v", sts.Int2, tt.wantInt2)
				}
				if sts.Limit != tt.wantLimit {
					t.Errorf("GetStats() Limit = %v, want %v", sts.Limit, tt.wantLimit)
				}
				if sts.Str1 != tt.wantStr1 {
					t.Errorf("GetStats() Str1 = %v, want %v", sts.Str1, tt.wantStr1)
				}
				if sts.Str2 != tt.wantStr2 {
					t.Errorf("GetStats() Str2 = %v, want %v", sts.Str2, tt.wantStr2)
				}
				if sts.Hits != tt.wantHits {
					t.Errorf("GetStats() Hits = %v, want %v", sts.Hits, tt.wantHits)
				}
			}
		})
	}
}

func TestStatsService_ResetStats(t *testing.T) {
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		repository func() adapters.StatsRepository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "successful reset",
			fields: fields{
				repository: func() adapters.StatsRepository {
					m := adapters.NewMockStatsRepository(ctrl)
					m.EXPECT().ResetStats().Return(nil).Times(1)
					return m
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatsService{
				repository: tt.fields.repository(),
			}
			if err := s.ResetStats(); (err != nil) != tt.wantErr {
				t.Errorf("ResetStats() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
