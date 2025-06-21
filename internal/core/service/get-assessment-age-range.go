package service

import (
	"errors"
	"sort"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (s *service) GetAssessmentAgeRange(req domain.GetAssessmentAgeRangeRequest) ([]domain.GetAssessmentAgeRangeResponse, error) {
	result, err := s.repo.GetAssessmentAgeRange(req)
	if err != nil {
		return nil, err
	}

	var filteredResult []domain.GetAssessmentAgeRangeResponse
	var correctOrder []string

	if req.AssessmentTypeId == "ASSMTT_1" {
		validAgeRangeIds := map[string]bool{
			"AGER_22": true, "AGER_23": true, "AGER_4": true, "AGER_5": true, "AGER_6": true,
			"AGER_7": true, "AGER_8": true, "AGER_9": true, "AGER_10": true, "AGER_11": true,
			"AGER_12": true, "AGER_13": true, "AGER_14": true, "AGER_15": true, "AGER_16": true,
			"AGER_17": true, "AGER_18": true, "AGER_19": true, "AGER_20": true, "AGER_21": true,
			"AGER_24": true, "AGER_25": true, "AGER_26": true,
		}
		for _, item := range result {
			if validAgeRangeIds[item.AgeRangeId] {
				item.AssessmentTypeId = &req.AssessmentTypeId
				item.AssessmentVideoUrl = item.AssessmentVideoUrlDspm
				filteredResult = append(filteredResult, item)
			}
		}

		correctOrder = []string{
			"AGER_22", "AGER_23", "AGER_4", "AGER_5", "AGER_6",
			"AGER_7", "AGER_8", "AGER_9", "AGER_10", "AGER_11",
			"AGER_12", "AGER_13", "AGER_14", "AGER_15", "AGER_16",
			"AGER_17", "AGER_18", "AGER_19", "AGER_20", "AGER_21",
			"AGER_24", "AGER_25", "AGER_26",
		}
	} else if req.AssessmentTypeId == "ASSMTT_2" {
		validAgeRangeIds := map[string]bool{
			"AGER_1": true, "AGER_2": true, "AGER_3": true, "AGER_4": true, "AGER_5": true,
			"AGER_6": true, "AGER_7": true, "AGER_8": true, "AGER_9": true, "AGER_10": true,
			"AGER_11": true, "AGER_12": true, "AGER_13": true, "AGER_14": true, "AGER_15": true,
			"AGER_16": true, "AGER_17": true, "AGER_18": true, "AGER_19": true, "AGER_20": true,
			"AGER_21": true,
		}
		for _, item := range result {
			if validAgeRangeIds[item.AgeRangeId] {
				item.AssessmentTypeId = &req.AssessmentTypeId
				item.AssessmentVideoUrl = item.AssessmentVideoUrlDaim // เอา DAIM มาใส่
				filteredResult = append(filteredResult, item)
			}
		}

		correctOrder = []string{
			"AGER_1", "AGER_2", "AGER_3", "AGER_4", "AGER_5",
			"AGER_6", "AGER_7", "AGER_8", "AGER_9", "AGER_10",
			"AGER_11", "AGER_12", "AGER_13", "AGER_14", "AGER_15",
			"AGER_16", "AGER_17", "AGER_18", "AGER_19", "AGER_20",
			"AGER_21",
		}
	} else {
		return nil, errors.New("invalid assessment type")
	}

	orderMap := make(map[string]int)
	for idx, id := range correctOrder {
		orderMap[id] = idx
	}

	sort.Slice(filteredResult, func(i, j int) bool {
		return orderMap[filteredResult[i].AgeRangeId] < orderMap[filteredResult[j].AgeRangeId]
	})

	return filteredResult, nil
}
