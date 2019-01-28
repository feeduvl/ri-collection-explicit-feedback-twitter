package main

import (
	"reflect"
	"testing"
	"time"
)

func TestTimeFrame_IsValid(t *testing.T) {
	type fields struct {
		since string
		until string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"identical dates", fields{since: "2018-05-03", until: "2018-05-03"}, false},
		{"since before until", fields{since: "2018-05-02", until: "2018-05-03"}, true},
		{"since after until", fields{since: "2018-05-03", until: "2018-05-02"}, false},
		{"contains random string", fields{since: "abcd", until: "2018-05-02"}, false},
		{"contains wrong date format", fields{since: "01.04.2018", until: "2018-05-02"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeFrame := TimeFrame{
				since: tt.fields.since,
				until: tt.fields.until,
			}
			if got := timeFrame.IsValid(); got != tt.want {
				t.Errorf("TimeFrame.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeFrameFromDays(t *testing.T) {
	type args struct {
		days int
	}
	tests := []struct {
		name string
		args args
		want TimeFrame
	}{
		{"since{yesterday}, until:{today}", args{days: 1}, TimeFrame{since: time.Now().Format(dateFormat), until: time.Now().AddDate(0, 0, 1).Format(dateFormat)}},
		{"negative number of days are not supported", args{days: -1}, TimeFrame{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeFrameFromDays(tt.args.days); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeFrameFromDays() = %v, want %v, name: %v", got, tt.want, tt.name)
			}
		})
	}
}

func TestTimeFrameFromSince(t *testing.T) {
	type args struct {
		since string
	}
	tests := []struct {
		name string
		args args
		want TimeFrame
	}{
		{"empty since", args{since: ""}, TimeFrame{}},
		{"correct format", args{since: "2018-05-3"}, TimeFrame{since: "2018-05-3", until: time.Now().Format(dateFormat)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeFrameFromSince(tt.args.since); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeFrameFromSince() = %v, want %v", got, tt.want)
			}
		})
	}
}
