package gothaibaht

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToText(t *testing.T) {
	var output string

	t.Run("should convert zero number to THB", func(t *testing.T) {
		output = ToText(0)
		assert.Equal(t, "ศูนย์บาทถ้วน", output)

		output = ToText(0.00)
		assert.Equal(t, "ศูนย์บาทถ้วน", output)
	})

	t.Run("should convert number to Satang", func(t *testing.T) {
		output = ToText(0.01)
		assert.Equal(t, "หนึ่งสตางค์", output)

		output = ToText(0.1)
		assert.Equal(t, "สิบสตางค์", output)

		output = ToText(0.10)
		assert.Equal(t, "สิบสตางค์", output)

		output = ToText(0.11)
		assert.Equal(t, "สิบเอ็ดสตางค์", output)

		output = ToText(0.12)
		assert.Equal(t, "สิบสองสตางค์", output)

		output = ToText(0.123)
		assert.Equal(t, "สิบสองสตางค์", output)

		output = ToText(0.2)
		assert.Equal(t, "ยี่สิบสตางค์", output)

		output = ToText(0.20)
		assert.Equal(t, "ยี่สิบสตางค์", output)

		output = ToText(0.21)
		assert.Equal(t, "ยี่สิบเอ็ดสตางค์", output)

		output = ToText(0.25)
		assert.Equal(t, "ยี่สิบห้าสตางค์", output)

		output = ToText(0.256)
		assert.Equal(t, "ยี่สิบห้าสตางค์", output)

		output = ToText(0.50)
		assert.Equal(t, "ห้าสิบสตางค์", output)

		output = ToText(0.75)
		assert.Equal(t, "เจ็ดสิบห้าสตางค์", output)

		output = ToText(0.99)
		assert.Equal(t, "เก้าสิบเก้าสตางค์", output)

		output = ToText(0.999)
		assert.Equal(t, "เก้าสิบเก้าสตางค์", output)
	})

	t.Run("should convert number to THB", func(t *testing.T) {
		output = ToText(1)
		assert.Equal(t, "หนึ่งบาทถ้วน", output)

		output = ToText(10)
		assert.Equal(t, "สิบบาทถ้วน", output)

		output = ToText(11)
		assert.Equal(t, "สิบเอ็ดบาทถ้วน", output)

		output = ToText(12)
		assert.Equal(t, "สิบสองบาทถ้วน", output)

		output = ToText(20)
		assert.Equal(t, "ยี่สิบบาทถ้วน", output)

		output = ToText(21)
		assert.Equal(t, "ยี่สิบเอ็ดบาทถ้วน", output)

		output = ToText(22)
		assert.Equal(t, "ยี่สิบสองบาทถ้วน", output)

		output = ToText(100)
		assert.Equal(t, "หนึ่งร้อยบาทถ้วน", output)

		output = ToText(101)
		assert.Equal(t, "หนึ่งร้อยเอ็ดบาทถ้วน", output)

		output = ToText(111)
		assert.Equal(t, "หนึ่งร้อยสิบเอ็ดบาทถ้วน", output)

		output = ToText(121)
		assert.Equal(t, "หนึ่งร้อยยี่สิบเอ็ดบาทถ้วน", output)
	})

	t.Run("should convert big number to THB", func(t *testing.T) {
		output = ToText(1000000)
		assert.Equal(t, "หนึ่งล้านบาทถ้วน", output)

		output = ToText(1000001)
		assert.Equal(t, "หนึ่งล้านเอ็ดบาทถ้วน", output)

		output = ToText(11000001)
		assert.Equal(t, "สิบเอ็ดล้านเอ็ดบาทถ้วน", output)

		output = ToText(11000000)
		assert.Equal(t, "สิบเอ็ดล้านบาทถ้วน", output)
	})

	t.Run("should convert multiple million to THB", func(t *testing.T) {
		output = ToText(1000000000000000000)
		assert.Equal(t, "หนึ่งล้านล้านล้านบาทถ้วน", output)

		output = ToText(1000000000001)
		assert.Equal(t, "หนึ่งล้านล้านเอ็ดบาทถ้วน", output)

		output = ToText(1000000000000)
		assert.Equal(t, "หนึ่งล้านล้านบาทถ้วน", output)

		output = ToText(1001000000001)
		assert.Equal(t, "หนึ่งล้านหนึ่งพันล้านเอ็ดบาทถ้วน", output)

		output = ToText(1001000001001)
		assert.Equal(t, "หนึ่งล้านหนึ่งพันล้านหนึ่งพันเอ็ดบาทถ้วน", output)

		output = ToText(1001000000000)
		assert.Equal(t, "หนึ่งล้านหนึ่งพันล้านบาทถ้วน", output)

		output = ToText(1000000000)
		assert.Equal(t, "หนึ่งพันล้านบาทถ้วน", output)

		output = ToText(10000000)
		assert.Equal(t, "สิบล้านบาทถ้วน", output)

		output = ToText(100000000)
		assert.Equal(t, "หนึ่งร้อยล้านบาทถ้วน", output)
	})

	t.Run("should convert complex number to THB", func(t *testing.T) {
		output = ToText(6321298)
		assert.Equal(t, "หกล้านสามแสนสองหมื่นหนึ่งพันสองร้อยเก้าสิบแปดบาทถ้วน", output)

		output = ToText(10034567)
		assert.Equal(t, "สิบล้านสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดบาทถ้วน", output)

		output = ToText(20034567)
		assert.Equal(t, "ยี่สิบล้านสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดบาทถ้วน", output)

		output = ToText(30034567.00)
		assert.Equal(t, "สามสิบล้านสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดบาทถ้วน", output)
	})

	t.Run("should convert number to THB with Satang", func(t *testing.T) {
		output = ToText(11.25)
		assert.Equal(t, "สิบเอ็ดบาทยี่สิบห้าสตางค์", output)

		output = ToText(100.50)
		assert.Equal(t, "หนึ่งร้อยบาทห้าสิบสตางค์", output)

		output = ToText(567.01)
		assert.Equal(t, "ห้าร้อยหกสิบเจ็ดบาทหนึ่งสตางค์", output)

		output = ToText(123456789.999)
		assert.Equal(t, "หนึ่งร้อยยี่สิบสามล้านสี่แสนห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทเก้าสิบเก้าสตางค์", output)
	})
}
