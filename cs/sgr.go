package cs

// Select Graphic Rendition (SGR) control sequences
const (
	SGR_0 = csi + "0" + m // SGR 0
	SGR_1 = csi + "1" + m // SGR 1
	SGR_2 = csi + "2" + m // SGR 2
	SGR_3 = csi + "3" + m // SGR 3
	SGR_4 = csi + "4" + m // SGR 4
	SGR_5 = csi + "5" + m // SGR 5
	SGR_6 = csi + "6" + m // SGR 6
	SGR_7 = csi + "7" + m // SGR 7
	SGR_8 = csi + "8" + m // SGR 8
	SGR_9 = csi + "9" + m // SGR 9

	SGR_10 = csi + "10" + m // SGR 10
	SGR_11 = csi + "11" + m // SGR 11
	SGR_12 = csi + "12" + m // SGR 12
	SGR_13 = csi + "13" + m // SGR 13
	SGR_14 = csi + "14" + m // SGR 14
	SGR_15 = csi + "15" + m // SGR 15
	SGR_16 = csi + "16" + m // SGR 16
	SGR_17 = csi + "17" + m // SGR 17
	SGR_18 = csi + "18" + m // SGR 18
	SGR_19 = csi + "19" + m // SGR 19

	SGR_20 = csi + "20" + m // SGR 20
	SGR_21 = csi + "21" + m // SGR 21
	SGR_22 = csi + "22" + m // SGR 22
	SGR_23 = csi + "23" + m // SGR 23
	SGR_24 = csi + "24" + m // SGR 24
	SGR_25 = csi + "25" + m // SGR 25
	SGR_26 = csi + "26" + m // SGR 26
	SGR_27 = csi + "27" + m // SGR 27
	SGR_28 = csi + "28" + m // SGR 28
	SGR_29 = csi + "29" + m // SGR 29

	SGR_30 = csi + "30" + m // SGR 30
	SGR_31 = csi + "31" + m // SGR 31
	SGR_32 = csi + "32" + m // SGR 32
	SGR_33 = csi + "33" + m // SGR 33
	SGR_34 = csi + "34" + m // SGR 34
	SGR_35 = csi + "35" + m // SGR 35
	SGR_36 = csi + "36" + m // SGR 36
	SGR_37 = csi + "37" + m // SGR 37
	SGR_38 = csi + "38" + m // SGR 38
	SGR_39 = csi + "39" + m // SGR 39

	SGR_40 = csi + "40" + m // SGR 40
	SGR_41 = csi + "41" + m // SGR 41
	SGR_42 = csi + "42" + m // SGR 42
	SGR_43 = csi + "43" + m // SGR 43
	SGR_44 = csi + "44" + m // SGR 44
	SGR_45 = csi + "45" + m // SGR 45
	SGR_46 = csi + "46" + m // SGR 46
	SGR_47 = csi + "47" + m // SGR 47
	SGR_48 = csi + "48" + m // SGR 48
	SGR_49 = csi + "49" + m // SGR 49

	SGR_50 = csi + "50" + m // SGR 50
	SGR_51 = csi + "51" + m // SGR 51
	SGR_52 = csi + "52" + m // SGR 52
	SGR_53 = csi + "53" + m // SGR 53
	SGR_54 = csi + "54" + m // SGR 54
	SGR_55 = csi + "55" + m // SGR 55
	SGR_56 = csi + "56" + m // SGR 56
	SGR_57 = csi + "57" + m // SGR 57
	SGR_58 = csi + "58" + m // SGR 58
	SGR_59 = csi + "59" + m // SGR 59

	SGR_60 = csi + "60" + m // SGR 60
	SGR_61 = csi + "61" + m // SGR 61
	SGR_62 = csi + "62" + m // SGR 62
	SGR_63 = csi + "63" + m // SGR 63
	SGR_64 = csi + "64" + m // SGR 64
	SGR_65 = csi + "65" + m // SGR 65
	SGR_66 = csi + "66" + m // SGR 66
	SGR_67 = csi + "67" + m // SGR 67
	SGR_68 = csi + "68" + m // SGR 68
	SGR_69 = csi + "69" + m // SGR 69

	SGR_70 = csi + "70" + m // SGR 70
	SGR_71 = csi + "71" + m // SGR 71
	SGR_72 = csi + "72" + m // SGR 72
	SGR_73 = csi + "73" + m // SGR 73
	SGR_74 = csi + "74" + m // SGR 74
	SGR_75 = csi + "75" + m // SGR 75
	SGR_76 = csi + "76" + m // SGR 76
	SGR_77 = csi + "77" + m // SGR 77
	SGR_78 = csi + "78" + m // SGR 78
	SGR_79 = csi + "79" + m // SGR 79

	SGR_80 = csi + "80" + m // SGR 80
	SGR_81 = csi + "81" + m // SGR 81
	SGR_82 = csi + "82" + m // SGR 82
	SGR_83 = csi + "83" + m // SGR 83
	SGR_84 = csi + "84" + m // SGR 84
	SGR_85 = csi + "85" + m // SGR 85
	SGR_86 = csi + "86" + m // SGR 86
	SGR_87 = csi + "87" + m // SGR 87
	SGR_88 = csi + "88" + m // SGR 88
	SGR_89 = csi + "89" + m // SGR 89

	SGR_90 = csi + "90" + m // SGR 90
	SGR_91 = csi + "91" + m // SGR 91
	SGR_92 = csi + "92" + m // SGR 92
	SGR_93 = csi + "93" + m // SGR 93
	SGR_94 = csi + "94" + m // SGR 94
	SGR_95 = csi + "95" + m // SGR 95
	SGR_96 = csi + "96" + m // SGR 96
	SGR_97 = csi + "97" + m // SGR 97
	SGR_98 = csi + "98" + m // SGR 98
	SGR_99 = csi + "99" + m // SGR 99

	SGR_100 = csi + "100" + m // SGR 100
	SGR_101 = csi + "101" + m // SGR 101
	SGR_102 = csi + "102" + m // SGR 102
	SGR_103 = csi + "103" + m // SGR 103
	SGR_104 = csi + "104" + m // SGR 104
	SGR_105 = csi + "105" + m // SGR 105
	SGR_106 = csi + "106" + m // SGR 106
	SGR_107 = csi + "107" + m // SGR 107
)
