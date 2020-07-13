package goconvey

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Add(i, j int) int {
	return i + j
}

func Division(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("除以 0")
	}
	return i / j, nil
}

func TestAdd(t *testing.T) {
	Convey("将两数相加", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func TestDivision(t *testing.T) {
	Convey("将两数相除", t, func() {

		Convey("除以非 0 数", func() {
			num, err := Division(10, 2)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, 5)
		})

		Convey("除以 0", func() {
			_, err := Division(10, 0)
			So(err, ShouldNotBeNil)
		})
	})
}
