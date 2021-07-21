package main

import (
	"testing"
)

func TestRectangle_Area(t *testing.T) {

	tests := []struct {
		name    string
		r       Rectangle
		want    float64
		wantErr bool
	}{
		{
			name: "should return ok if all right",
			r: Rectangle{
				height: 2,
				width:  2,
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "should return error if unavailable value ",
			r: Rectangle{
				height: 0,
				width:  2,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name    string
		c       Circle
		want    float64
		wantErr bool
	}{
		{
			name: "should return ok if all right",
			c: Circle{
				radius: 10,
			},
			want:    314.1592653589793,
			wantErr: false,
		},
		{
			name: "should return error if unavailable value",
			c: Circle{
				radius: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		r       Rectangle
		want    float64
		wantErr bool
	}{
		{
			name: "should return ok if all right",
			r: Rectangle{
				height: 2,
				width:  2,
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "should return error if unavailable value ",
			r: Rectangle{
				height: 0,
				width:  2,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		c       Circle
		want    float64
		wantErr bool
	}{
		{
			name: "should return ok if all right",
			c: Circle{
				radius: 10,
			},
			want:    62.83185307179586,
			wantErr: false,
		},
		{
			name: "should return error if unavailable value",
			c: Circle{
				radius: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Circle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
