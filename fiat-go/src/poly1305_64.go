/* Autogenerated: src/ExtractionOCaml/unsaturated_solinas --lang=Go --no-wide-int --cmovznz-by-mul --widen-carry --widen-bytes poly1305 3 '2^130 - 5' 64 carry_mul carry_square carry add sub opp selectznz to_bytes from_bytes */
/* curve description: poly1305 */
/* requested operations: carry_mul, carry_square, carry, add, sub, opp, selectznz, to_bytes, from_bytes */
/* n = 3 (from "3") */
/* s-c = 2^130 - [(1, 5)] (from "2^130 - 5") */
/* machine_wordsize = 64 (from "64") */

/* Computed values: */
/* carry_chain = [0, 1, 2, 0, 1] */

package fiat_poly1305

import "math/bits"


/*
 * The function fiat_poly1305_addcarryx_u44 is an addition with carry.
 * Postconditions:
 *   out1 = (arg1 + arg2 + arg3) mod 2^44
 *   out2 = ⌊(arg1 + arg2 + arg3) / 2^44⌋
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0xfffffffffff]
 *   arg3: [0x0 ~> 0xfffffffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0xfffffffffff]
 *   out2: [0x0 ~> 0x1]
 */
/*inline*/
func fiat_poly1305_addcarryx_u44(out1 *uint64, out2 *uint64, arg1 uint64, arg2 uint64, arg3 uint64) {
  var x1 uint64 = ((arg1 + arg2) + arg3)
  var x2 uint64 = (x1 & 0xfffffffffff)
  var x3 uint64 = (x1 >> 44)
  *out1 = x2
  *out2 = x3
}

/*
 * The function fiat_poly1305_subborrowx_u44 is a subtraction with borrow.
 * Postconditions:
 *   out1 = (-arg1 + arg2 + -arg3) mod 2^44
 *   out2 = -⌊(-arg1 + arg2 + -arg3) / 2^44⌋
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0xfffffffffff]
 *   arg3: [0x0 ~> 0xfffffffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0xfffffffffff]
 *   out2: [0x0 ~> 0x1]
 */
/*inline*/
func fiat_poly1305_subborrowx_u44(out1 *uint64, out2 *uint64, arg1 uint64, arg2 uint64, arg3 uint64) {
  var x1 int64 = ((int64(arg2) - int64(arg1)) - int64(arg3))
  var x2 int64 = (x1 >> 44)
  var x3 uint64 = (uint64(x1) & 0xfffffffffff)
  *out1 = x3
  *out2 = (uint64(0x0) - uint64(x2))
}

/*
 * The function fiat_poly1305_addcarryx_u43 is an addition with carry.
 * Postconditions:
 *   out1 = (arg1 + arg2 + arg3) mod 2^43
 *   out2 = ⌊(arg1 + arg2 + arg3) / 2^43⌋
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0x7ffffffffff]
 *   arg3: [0x0 ~> 0x7ffffffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0x7ffffffffff]
 *   out2: [0x0 ~> 0x1]
 */
/*inline*/
func fiat_poly1305_addcarryx_u43(out1 *uint64, out2 *uint64, arg1 uint64, arg2 uint64, arg3 uint64) {
  var x1 uint64 = ((arg1 + arg2) + arg3)
  var x2 uint64 = (x1 & 0x7ffffffffff)
  var x3 uint64 = (x1 >> 43)
  *out1 = x2
  *out2 = x3
}

/*
 * The function fiat_poly1305_subborrowx_u43 is a subtraction with borrow.
 * Postconditions:
 *   out1 = (-arg1 + arg2 + -arg3) mod 2^43
 *   out2 = -⌊(-arg1 + arg2 + -arg3) / 2^43⌋
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0x7ffffffffff]
 *   arg3: [0x0 ~> 0x7ffffffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0x7ffffffffff]
 *   out2: [0x0 ~> 0x1]
 */
/*inline*/
func fiat_poly1305_subborrowx_u43(out1 *uint64, out2 *uint64, arg1 uint64, arg2 uint64, arg3 uint64) {
  var x1 int64 = ((int64(arg2) - int64(arg1)) - int64(arg3))
  var x2 int64 = (x1 >> 43)
  var x3 uint64 = (uint64(x1) & 0x7ffffffffff)
  *out1 = x3
  *out2 = (uint64(0x0) - uint64(x2))
}

/*
 * The function fiat_poly1305_cmovznz_u64 is a single-word conditional move.
 * Postconditions:
 *   out1 = (if arg1 = 0 then arg2 else arg3)
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0xffffffffffffffff]
 *   arg3: [0x0 ~> 0xffffffffffffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0xffffffffffffffff]
 */
/*inline*/
func fiat_poly1305_cmovznz_u64(out1 *uint64, arg1 uint64, arg2 uint64, arg3 uint64) {
  var x1 uint64 = (arg1 * 0xffffffffffffffff)
  var x2 uint64 = ((x1 & arg3) | ((^x1) & arg2))
  *out1 = x2
}

/*
 * The function fiat_poly1305_carry_mul multiplies two field elements and reduces the result.
 * Postconditions:
 *   eval out1 mod m = (eval arg1 * eval arg2) mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0x34cccccccccb], [0x0 ~> 0x1a6666666664], [0x0 ~> 0x1a6666666664]]
 *   arg2: [[0x0 ~> 0x34cccccccccb], [0x0 ~> 0x1a6666666664], [0x0 ~> 0x1a6666666664]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x119999999999], [0x0 ~> 0x8cccccccccc], [0x0 ~> 0x8cccccccccc]]
 */
/*inline*/
func fiat_poly1305_carry_mul(out1 *[3]uint64, arg1 *[3]uint64, arg2 *[3]uint64) {
  var x1 uint64
  var x2 uint64
  x1, x2 = bits.Mul64((arg1[2]), ((arg2[2]) * 0x5))
  var x3 uint64
  var x4 uint64
  x3, x4 = bits.Mul64((arg1[2]), ((arg2[1]) * 0xa))
  var x5 uint64
  var x6 uint64
  x5, x6 = bits.Mul64((arg1[1]), ((arg2[2]) * 0xa))
  var x7 uint64
  var x8 uint64
  x7, x8 = bits.Mul64((arg1[2]), (arg2[0]))
  var x9 uint64
  var x10 uint64
  x9, x10 = bits.Mul64((arg1[1]), ((arg2[1]) * 0x2))
  var x11 uint64
  var x12 uint64
  x11, x12 = bits.Mul64((arg1[1]), (arg2[0]))
  var x13 uint64
  var x14 uint64
  x13, x14 = bits.Mul64((arg1[0]), (arg2[2]))
  var x15 uint64
  var x16 uint64
  x15, x16 = bits.Mul64((arg1[0]), (arg2[1]))
  var x17 uint64
  var x18 uint64
  x17, x18 = bits.Mul64((arg1[0]), (arg2[0]))
  var x19 uint64
  var x20 uint64
  x19, x20 = bits.Add64(x5, x3, 0x0)
  var x21 uint64
  x21, _ = bits.Add64(x6, x4, x20)
  var x23 uint64
  var x24 uint64
  x23, x24 = bits.Add64(x17, x19, 0x0)
  var x25 uint64
  x25, _ = bits.Add64(x18, x21, x24)
  var x27 uint64 = ((x23 >> 44) | ((x25 << 20) & 0xffffffffffffffff))
  var x28 uint64 = (x23 & 0xfffffffffff)
  var x29 uint64
  var x30 uint64
  x29, x30 = bits.Add64(x9, x7, 0x0)
  var x31 uint64
  x31, _ = bits.Add64(x10, x8, x30)
  var x33 uint64
  var x34 uint64
  x33, x34 = bits.Add64(x13, x29, 0x0)
  var x35 uint64
  x35, _ = bits.Add64(x14, x31, x34)
  var x37 uint64
  var x38 uint64
  x37, x38 = bits.Add64(x11, x1, 0x0)
  var x39 uint64
  x39, _ = bits.Add64(x12, x2, x38)
  var x41 uint64
  var x42 uint64
  x41, x42 = bits.Add64(x15, x37, 0x0)
  var x43 uint64
  x43, _ = bits.Add64(x16, x39, x42)
  var x45 uint64
  var x46 uint64
  x45, x46 = bits.Add64(x27, x41, 0x0)
  var x47 uint64 = (x46 + x43)
  var x48 uint64 = ((x45 >> 43) | ((x47 << 21) & 0xffffffffffffffff))
  var x49 uint64 = (x45 & 0x7ffffffffff)
  var x50 uint64
  var x51 uint64
  x50, x51 = bits.Add64(x48, x33, 0x0)
  var x52 uint64 = (x51 + x35)
  var x53 uint64 = ((x50 >> 43) | ((x52 << 21) & 0xffffffffffffffff))
  var x54 uint64 = (x50 & 0x7ffffffffff)
  var x55 uint64 = (x53 * 0x5)
  var x56 uint64 = (x28 + x55)
  var x57 uint64 = (x56 >> 44)
  var x58 uint64 = (x56 & 0xfffffffffff)
  var x59 uint64 = (x57 + x49)
  var x60 uint64 = (x59 >> 43)
  var x61 uint64 = (x59 & 0x7ffffffffff)
  var x62 uint64 = (x60 + x54)
  out1[0] = x58
  out1[1] = x61
  out1[2] = x62
}

/*
 * The function fiat_poly1305_carry_square squares a field element and reduces the result.
 * Postconditions:
 *   eval out1 mod m = (eval arg1 * eval arg1) mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0x34cccccccccb], [0x0 ~> 0x1a6666666664], [0x0 ~> 0x1a6666666664]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x119999999999], [0x0 ~> 0x8cccccccccc], [0x0 ~> 0x8cccccccccc]]
 */
/*inline*/
func fiat_poly1305_carry_square(out1 *[3]uint64, arg1 *[3]uint64) {
  var x1 uint64 = ((arg1[2]) * 0x5)
  var x2 uint64 = (x1 * 0x2)
  var x3 uint64 = ((arg1[2]) * 0x2)
  var x4 uint64 = ((arg1[1]) * 0x2)
  var x5 uint64
  var x6 uint64
  x5, x6 = bits.Mul64((arg1[2]), x1)
  var x7 uint64
  var x8 uint64
  x7, x8 = bits.Mul64((arg1[1]), (x2 * 0x2))
  var x9 uint64
  var x10 uint64
  x9, x10 = bits.Mul64((arg1[1]), ((arg1[1]) * 0x2))
  var x11 uint64
  var x12 uint64
  x11, x12 = bits.Mul64((arg1[0]), x3)
  var x13 uint64
  var x14 uint64
  x13, x14 = bits.Mul64((arg1[0]), x4)
  var x15 uint64
  var x16 uint64
  x15, x16 = bits.Mul64((arg1[0]), (arg1[0]))
  var x17 uint64
  var x18 uint64
  x17, x18 = bits.Add64(x15, x7, 0x0)
  var x19 uint64
  x19, _ = bits.Add64(x16, x8, x18)
  var x21 uint64 = ((x17 >> 44) | ((x19 << 20) & 0xffffffffffffffff))
  var x22 uint64 = (x17 & 0xfffffffffff)
  var x23 uint64
  var x24 uint64
  x23, x24 = bits.Add64(x11, x9, 0x0)
  var x25 uint64
  x25, _ = bits.Add64(x12, x10, x24)
  var x27 uint64
  var x28 uint64
  x27, x28 = bits.Add64(x13, x5, 0x0)
  var x29 uint64
  x29, _ = bits.Add64(x14, x6, x28)
  var x31 uint64
  var x32 uint64
  x31, x32 = bits.Add64(x21, x27, 0x0)
  var x33 uint64 = (x32 + x29)
  var x34 uint64 = ((x31 >> 43) | ((x33 << 21) & 0xffffffffffffffff))
  var x35 uint64 = (x31 & 0x7ffffffffff)
  var x36 uint64
  var x37 uint64
  x36, x37 = bits.Add64(x34, x23, 0x0)
  var x38 uint64 = (x37 + x25)
  var x39 uint64 = ((x36 >> 43) | ((x38 << 21) & 0xffffffffffffffff))
  var x40 uint64 = (x36 & 0x7ffffffffff)
  var x41 uint64 = (x39 * 0x5)
  var x42 uint64 = (x22 + x41)
  var x43 uint64 = (x42 >> 44)
  var x44 uint64 = (x42 & 0xfffffffffff)
  var x45 uint64 = (x43 + x35)
  var x46 uint64 = (x45 >> 43)
  var x47 uint64 = (x45 & 0x7ffffffffff)
  var x48 uint64 = (x46 + x40)
  out1[0] = x44
  out1[1] = x47
  out1[2] = x48
}

/*
 * The function fiat_poly1305_carry reduces a field element.
 * Postconditions:
 *   eval out1 mod m = eval arg1 mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0x34cccccccccb], [0x0 ~> 0x1a6666666664], [0x0 ~> 0x1a6666666664]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x119999999999], [0x0 ~> 0x8cccccccccc], [0x0 ~> 0x8cccccccccc]]
 */
/*inline*/
func fiat_poly1305_carry(out1 *[3]uint64, arg1 *[3]uint64) {
  var x1 uint64 = (arg1[0])
  var x2 uint64 = ((x1 >> 44) + (arg1[1]))
  var x3 uint64 = ((x2 >> 43) + (arg1[2]))
  var x4 uint64 = ((x1 & 0xfffffffffff) + ((x3 >> 43) * 0x5))
  var x5 uint64 = ((x4 >> 44) + (x2 & 0x7ffffffffff))
  var x6 uint64 = (x4 & 0xfffffffffff)
  var x7 uint64 = (x5 & 0x7ffffffffff)
  var x8 uint64 = ((x5 >> 43) + (x3 & 0x7ffffffffff))
  out1[0] = x6
  out1[1] = x7
  out1[2] = x8
}

/*
 * The function fiat_poly1305_add adds two field elements.
 * Postconditions:
 *   eval out1 mod m = (eval arg1 + eval arg2) mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0x119999999999], [0x0 ~> 0x8cccccccccc], [0x0 ~> 0x8cccccccccc]]
 *   arg2: [[0x0 ~> 0x119999999999], [0x0 ~> 0x8cccccccccc], [0x0 ~> 0x8cccccccccc]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x34cccccccccb], [0x0 ~> 0x1a6666666664], [0x0 ~> 0x1a6666666664]]
 */
/*inline*/
func fiat_poly1305_add(out1 *[3]uint64, arg1 *[3]uint64, arg2 *[3]uint64) {
  var x1 uint64 = ((arg1[0]) + (arg2[0]))
  var x2 uint64 = ((arg1[1]) + (arg2[1]))
  var x3 uint64 = ((arg1[2]) + (arg2[2]))
  out1[0] = x1
  out1[1] = x2
  out1[2] = x3
}

/*
 * The function fiat_poly1305_sub subtracts two field elements.
 * Postconditions:
 *   eval out1 mod m = (eval arg1 - eval arg2) mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0x119999999999], [0x0 ~> 0x8cccccccccc], [0x0 ~> 0x8cccccccccc]]
 *   arg2: [[0x0 ~> 0x119999999999], [0x0 ~> 0x8cccccccccc], [0x0 ~> 0x8cccccccccc]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x34cccccccccb], [0x0 ~> 0x1a6666666664], [0x0 ~> 0x1a6666666664]]
 */
/*inline*/
func fiat_poly1305_sub(out1 *[3]uint64, arg1 *[3]uint64, arg2 *[3]uint64) {
  var x1 uint64 = ((0x1ffffffffff6 + (arg1[0])) - (arg2[0]))
  var x2 uint64 = ((0xffffffffffe + (arg1[1])) - (arg2[1]))
  var x3 uint64 = ((0xffffffffffe + (arg1[2])) - (arg2[2]))
  out1[0] = x1
  out1[1] = x2
  out1[2] = x3
}

/*
 * The function fiat_poly1305_opp negates a field element.
 * Postconditions:
 *   eval out1 mod m = -eval arg1 mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0x119999999999], [0x0 ~> 0x8cccccccccc], [0x0 ~> 0x8cccccccccc]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x34cccccccccb], [0x0 ~> 0x1a6666666664], [0x0 ~> 0x1a6666666664]]
 */
/*inline*/
func fiat_poly1305_opp(out1 *[3]uint64, arg1 *[3]uint64) {
  var x1 uint64 = (0x1ffffffffff6 - (arg1[0]))
  var x2 uint64 = (0xffffffffffe - (arg1[1]))
  var x3 uint64 = (0xffffffffffe - (arg1[2]))
  out1[0] = x1
  out1[1] = x2
  out1[2] = x3
}

/*
 * The function fiat_poly1305_selectznz is a multi-limb conditional select.
 * Postconditions:
 *   eval out1 = (if arg1 = 0 then eval arg2 else eval arg3)
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 *   arg3: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 */
/*inline*/
func fiat_poly1305_selectznz(out1 *[3]uint64, arg1 uint64, arg2 *[3]uint64, arg3 *[3]uint64) {
  var x1 uint64
  fiat_poly1305_cmovznz_u64(&x1, arg1, (arg2[0]), (arg3[0]))
  var x2 uint64
  fiat_poly1305_cmovznz_u64(&x2, arg1, (arg2[1]), (arg3[1]))
  var x3 uint64
  fiat_poly1305_cmovznz_u64(&x3, arg1, (arg2[2]), (arg3[2]))
  out1[0] = x1
  out1[1] = x2
  out1[2] = x3
}

/*
 * The function fiat_poly1305_to_bytes serializes a field element to bytes in little-endian order.
 * Postconditions:
 *   out1 = map (λ x, ⌊((eval arg1 mod m) mod 2^(8 * (x + 1))) / 2^(8 * x)⌋) [0..16]
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0x119999999999], [0x0 ~> 0x8cccccccccc], [0x0 ~> 0x8cccccccccc]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x3]]
 */
/*inline*/
func fiat_poly1305_to_bytes(out1 *[17]uint64, arg1 *[3]uint64) {
  var x1 uint64
  var x2 uint64
  fiat_poly1305_subborrowx_u44(&x1, &x2, uint64(0x0), (arg1[0]), 0xffffffffffb)
  var x3 uint64
  var x4 uint64
  fiat_poly1305_subborrowx_u43(&x3, &x4, x2, (arg1[1]), 0x7ffffffffff)
  var x5 uint64
  var x6 uint64
  fiat_poly1305_subborrowx_u43(&x5, &x6, x4, (arg1[2]), 0x7ffffffffff)
  var x7 uint64
  fiat_poly1305_cmovznz_u64(&x7, x6, uint64(0x0), 0xffffffffffffffff)
  var x8 uint64
  var x9 uint64
  fiat_poly1305_addcarryx_u44(&x8, &x9, 0x0, x1, (x7 & 0xffffffffffb))
  var x10 uint64
  var x11 uint64
  fiat_poly1305_addcarryx_u43(&x10, &x11, x9, x3, (x7 & 0x7ffffffffff))
  var x12 uint64
  var x13 uint64
  fiat_poly1305_addcarryx_u43(&x12, &x13, x11, x5, (x7 & 0x7ffffffffff))
  var x14 uint64 = (x12 << 7)
  var x15 uint64 = (x10 << 4)
  var x16 uint64 = (x8 >> 8)
  var x17 uint64 = (x8 & 0xff)
  var x18 uint64 = (x16 >> 8)
  var x19 uint64 = (x16 & 0xff)
  var x20 uint64 = (x18 >> 8)
  var x21 uint64 = (x18 & 0xff)
  var x22 uint64 = (x20 >> 8)
  var x23 uint64 = (x20 & 0xff)
  var x24 uint64 = (x22 >> 8)
  var x25 uint64 = (x22 & 0xff)
  var x26 uint64 = (x24 + x15)
  var x27 uint64 = (x26 >> 8)
  var x28 uint64 = (x26 & 0xff)
  var x29 uint64 = (x27 >> 8)
  var x30 uint64 = (x27 & 0xff)
  var x31 uint64 = (x29 >> 8)
  var x32 uint64 = (x29 & 0xff)
  var x33 uint64 = (x31 >> 8)
  var x34 uint64 = (x31 & 0xff)
  var x35 uint64 = (x33 >> 8)
  var x36 uint64 = (x33 & 0xff)
  var x37 uint64 = (x35 + x14)
  var x38 uint64 = (x37 >> 8)
  var x39 uint64 = (x37 & 0xff)
  var x40 uint64 = (x38 >> 8)
  var x41 uint64 = (x38 & 0xff)
  var x42 uint64 = (x40 >> 8)
  var x43 uint64 = (x40 & 0xff)
  var x44 uint64 = (x42 >> 8)
  var x45 uint64 = (x42 & 0xff)
  var x46 uint64 = (x44 >> 8)
  var x47 uint64 = (x44 & 0xff)
  var x48 uint64 = (x46 >> 8)
  var x49 uint64 = (x46 & 0xff)
  out1[0] = x17
  out1[1] = x19
  out1[2] = x21
  out1[3] = x23
  out1[4] = x25
  out1[5] = x28
  out1[6] = x30
  out1[7] = x32
  out1[8] = x34
  out1[9] = x36
  out1[10] = x39
  out1[11] = x41
  out1[12] = x43
  out1[13] = x45
  out1[14] = x47
  out1[15] = x49
  out1[16] = x48
}

/*
 * The function fiat_poly1305_from_bytes deserializes a field element from bytes in little-endian order.
 * Postconditions:
 *   eval out1 mod m = bytes_eval arg1 mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x3]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x119999999999], [0x0 ~> 0x8cccccccccc], [0x0 ~> 0x8cccccccccc]]
 */
/*inline*/
func fiat_poly1305_from_bytes(out1 *[3]uint64, arg1 *[17]uint64) {
  var x1 uint64 = ((arg1[16]) << 41)
  var x2 uint64 = ((arg1[15]) << 33)
  var x3 uint64 = ((arg1[14]) << 25)
  var x4 uint64 = ((arg1[13]) << 17)
  var x5 uint64 = ((arg1[12]) << 9)
  var x6 uint64 = ((arg1[11]) * 0x2)
  var x7 uint64 = ((arg1[10]) << 36)
  var x8 uint64 = ((arg1[9]) << 28)
  var x9 uint64 = ((arg1[8]) << 20)
  var x10 uint64 = ((arg1[7]) << 12)
  var x11 uint64 = ((arg1[6]) << 4)
  var x12 uint64 = ((arg1[5]) << 40)
  var x13 uint64 = ((arg1[4]) << 32)
  var x14 uint64 = ((arg1[3]) << 24)
  var x15 uint64 = ((arg1[2]) << 16)
  var x16 uint64 = ((arg1[1]) << 8)
  var x17 uint64 = (arg1[0])
  var x18 uint64 = (x17 + (x16 + (x15 + (x14 + (x13 + x12)))))
  var x19 uint64 = (x18 >> 44)
  var x20 uint64 = (x18 & 0xfffffffffff)
  var x21 uint64 = (x6 + (x5 + (x4 + (x3 + (x2 + x1)))))
  var x22 uint64 = (x11 + (x10 + (x9 + (x8 + x7))))
  var x23 uint64 = (x19 + x22)
  var x24 uint64 = (x23 >> 43)
  var x25 uint64 = (x23 & 0x7ffffffffff)
  var x26 uint64 = (x24 + x21)
  out1[0] = x20
  out1[1] = x25
  out1[2] = x26
}

