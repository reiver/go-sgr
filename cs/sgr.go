package cs

// Select Graphic Rendition (SGR) control sequences
const (
	SGR_0 Code = csi + "0" + m // SGR 0
	SGR_1 Code = csi + "1" + m // SGR 1
	SGR_2 Code = csi + "2" + m // SGR 2
	SGR_3 Code = csi + "3" + m // SGR 3
	SGR_4 Code = csi + "4" + m // SGR 4
	SGR_5 Code = csi + "5" + m // SGR 5
	SGR_6 Code = csi + "6" + m // SGR 6
	SGR_7 Code = csi + "7" + m // SGR 7
	SGR_8 Code = csi + "8" + m // SGR 8
	SGR_9 Code = csi + "9" + m // SGR 9

	SGR_10 Code = csi + "10" + m // SGR 10
	SGR_11 Code = csi + "11" + m // SGR 11
	SGR_12 Code = csi + "12" + m // SGR 12
	SGR_13 Code = csi + "13" + m // SGR 13
	SGR_14 Code = csi + "14" + m // SGR 14
	SGR_15 Code = csi + "15" + m // SGR 15
	SGR_16 Code = csi + "16" + m // SGR 16
	SGR_17 Code = csi + "17" + m // SGR 17
	SGR_18 Code = csi + "18" + m // SGR 18
	SGR_19 Code = csi + "19" + m // SGR 19

	SGR_20 Code = csi + "20" + m // SGR 20
	SGR_21 Code = csi + "21" + m // SGR 21
	SGR_22 Code = csi + "22" + m // SGR 22
	SGR_23 Code = csi + "23" + m // SGR 23
	SGR_24 Code = csi + "24" + m // SGR 24
	SGR_25 Code = csi + "25" + m // SGR 25
	SGR_26 Code = csi + "26" + m // SGR 26
	SGR_27 Code = csi + "27" + m // SGR 27
	SGR_28 Code = csi + "28" + m // SGR 28
	SGR_29 Code = csi + "29" + m // SGR 29

	SGR_30 Code = csi + "30" + m // SGR 30
	SGR_31 Code = csi + "31" + m // SGR 31
	SGR_32 Code = csi + "32" + m // SGR 32
	SGR_33 Code = csi + "33" + m // SGR 33
	SGR_34 Code = csi + "34" + m // SGR 34
	SGR_35 Code = csi + "35" + m // SGR 35
	SGR_36 Code = csi + "36" + m // SGR 36
	SGR_37 Code = csi + "37" + m // SGR 37
	SGR_38 Code = csi + "38" + m // SGR 38
	SGR_39 Code = csi + "39" + m // SGR 39

	SGR_40 Code = csi + "40" + m // SGR 40
	SGR_41 Code = csi + "41" + m // SGR 41
	SGR_42 Code = csi + "42" + m // SGR 42
	SGR_43 Code = csi + "43" + m // SGR 43
	SGR_44 Code = csi + "44" + m // SGR 44
	SGR_45 Code = csi + "45" + m // SGR 45
	SGR_46 Code = csi + "46" + m // SGR 46
	SGR_47 Code = csi + "47" + m // SGR 47
	SGR_48 Code = csi + "48" + m // SGR 48
	SGR_49 Code = csi + "49" + m // SGR 49

	SGR_50 Code = csi + "50" + m // SGR 50
	SGR_51 Code = csi + "51" + m // SGR 51
	SGR_52 Code = csi + "52" + m // SGR 52
	SGR_53 Code = csi + "53" + m // SGR 53
	SGR_54 Code = csi + "54" + m // SGR 54
	SGR_55 Code = csi + "55" + m // SGR 55
	SGR_56 Code = csi + "56" + m // SGR 56
	SGR_57 Code = csi + "57" + m // SGR 57
	SGR_58 Code = csi + "58" + m // SGR 58
	SGR_59 Code = csi + "59" + m // SGR 59

	SGR_60 Code = csi + "60" + m // SGR 60
	SGR_61 Code = csi + "61" + m // SGR 61
	SGR_62 Code = csi + "62" + m // SGR 62
	SGR_63 Code = csi + "63" + m // SGR 63
	SGR_64 Code = csi + "64" + m // SGR 64
	SGR_65 Code = csi + "65" + m // SGR 65
	SGR_66 Code = csi + "66" + m // SGR 66
	SGR_67 Code = csi + "67" + m // SGR 67
	SGR_68 Code = csi + "68" + m // SGR 68
	SGR_69 Code = csi + "69" + m // SGR 69

	SGR_70 Code = csi + "70" + m // SGR 70
	SGR_71 Code = csi + "71" + m // SGR 71
	SGR_72 Code = csi + "72" + m // SGR 72
	SGR_73 Code = csi + "73" + m // SGR 73
	SGR_74 Code = csi + "74" + m // SGR 74
	SGR_75 Code = csi + "75" + m // SGR 75
	SGR_76 Code = csi + "76" + m // SGR 76
	SGR_77 Code = csi + "77" + m // SGR 77
	SGR_78 Code = csi + "78" + m // SGR 78
	SGR_79 Code = csi + "79" + m // SGR 79

	SGR_80 Code = csi + "80" + m // SGR 80
	SGR_81 Code = csi + "81" + m // SGR 81
	SGR_82 Code = csi + "82" + m // SGR 82
	SGR_83 Code = csi + "83" + m // SGR 83
	SGR_84 Code = csi + "84" + m // SGR 84
	SGR_85 Code = csi + "85" + m // SGR 85
	SGR_86 Code = csi + "86" + m // SGR 86
	SGR_87 Code = csi + "87" + m // SGR 87
	SGR_88 Code = csi + "88" + m // SGR 88
	SGR_89 Code = csi + "89" + m // SGR 89

	SGR_90 Code = csi + "90" + m // SGR 90
	SGR_91 Code = csi + "91" + m // SGR 91
	SGR_92 Code = csi + "92" + m // SGR 92
	SGR_93 Code = csi + "93" + m // SGR 93
	SGR_94 Code = csi + "94" + m // SGR 94
	SGR_95 Code = csi + "95" + m // SGR 95
	SGR_96 Code = csi + "96" + m // SGR 96
	SGR_97 Code = csi + "97" + m // SGR 97
	SGR_98 Code = csi + "98" + m // SGR 98
	SGR_99 Code = csi + "99" + m // SGR 99

	SGR_100 Code = csi + "100" + m // SGR 100
	SGR_101 Code = csi + "101" + m // SGR 101
	SGR_102 Code = csi + "102" + m // SGR 102
	SGR_103 Code = csi + "103" + m // SGR 103
	SGR_104 Code = csi + "104" + m // SGR 104
	SGR_105 Code = csi + "105" + m // SGR 105
	SGR_106 Code = csi + "106" + m // SGR 106
	SGR_107 Code = csi + "107" + m // SGR 107
)
