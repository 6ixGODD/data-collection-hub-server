package service_test

import (
	"testing"
)

func TestGetInstructionData(t *testing.T) {
	resp, err := adminDataAuditService.GetInstructionData(ctx, mockInstructionData.RandomInstructionDataID())

}
