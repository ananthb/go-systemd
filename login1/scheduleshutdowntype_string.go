// Code generated by "stringer -type=ScheduleShutdownType -linecomment"; DO NOT EDIT.

package login1

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ScheduleTypePowerOff-0]
	_ = x[ScheduleTypeDryPowerOff-1]
	_ = x[ScheduleTypeReboot-2]
	_ = x[ScheduleTypeDryReboot-3]
	_ = x[ScheduleTypeHalt-4]
	_ = x[SceduleTypeDryHalt-5]
}

const _ScheduleShutdownType_name = "poweroffdry-poweroffrebootdry-reboothaltdry-halt"

var _ScheduleShutdownType_index = [...]uint8{0, 8, 20, 26, 36, 40, 48}

func (i ScheduleShutdownType) String() string {
	if i < 0 || i >= ScheduleShutdownType(len(_ScheduleShutdownType_index)-1) {
		return "ScheduleShutdownType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ScheduleShutdownType_name[_ScheduleShutdownType_index[i]:_ScheduleShutdownType_index[i+1]]
}
