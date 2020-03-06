package main

const defaultPartDataJson = `{
	"box": {
		"x": 1,
		"y": 1,
		"z": 1
	},
	"color": "df7f00",
	"density": 250,
	"physicsMaterial": "Metal",
	"qualityLevel": 3,
	"renderable": {
		"lodList": [
			{
				"mesh": "$GAME_DATA/Objects/Mesh/interactive/obj_interactive_logicgate_off.fbx",
				"pose0": "$GAME_DATA/Objects/Mesh/interactive/obj_interactive_logicgate_on.fbx",
				"subMeshList": [
					{
						"material": "PoseAnimDifAsgNor",
						"textureList": [
							"$GAME_DATA/Objects/Textures/interactive/obj_interactive_logicgate_dif.tga",
							"$GAME_DATA/Objects/Textures/interactive/obj_interactive_logicgate_asg.tga",
							"$GAME_DATA/Objects/Textures/interactive/obj_interactive_logicgate_nor.tga"
						]
					},
					{
						"custom": {
							"uv0": {
								"u": 0.16666667,
								"v": 0.1645
							}
						},
						"material": "UVAnimDifAsgNor",
						"textureList": [
							"$GAME_DATA/Objects/Textures/interactive/obj_interactive_logicgate_dif.tga",
							"$GAME_DATA/Objects/Textures/interactive/obj_interactive_logicgate_asg.tga",
							"$GAME_DATA/Objects/Textures/interactive/obj_interactive_logicgate_nor.tga"
						]
					}
				]
			}
		]
	},
	"rotationSet": "PropYZ",
	"scripted": {
		"classname": "LogicExample",
		"filename": "$MOD_DATA/Scripts/Script.lua"
	}
}`

var smcolors = [][]string{
	[]string{
		"eeeeee",
		"7f7f7f",
		"4a4a4a",
		"222222",
	},
	[]string{
		"f5f071",
		"e2db13",
		"817c00",
		"323000",
	},
	[]string{
		"cbf66f",
		"a0ea00",
		"577d07",
		"375000",
	},
	[]string{
		"68ff88",
		"19e753",
		"0e8031",
		"064023",
	},
	[]string{
		"7eeded",
		"2ce6e6",
		"118787",
		"0a4444",
	},
	[]string{
		"4c6fe3",
		"0a3ee2",
		"0f2e91",
		"0a1d5a",
	},
	[]string{
		"ae79f0",
		"7514ed",
		"500aa6",
		"35086c",
	},
	[]string{
		"ee7bf0",
		"cf11d2",
		"720a74",
		"520653",
	},
	[]string{
		"f06767",
		"d02525",
		"7c0000",
		"560202",
	},
	[]string{
		"eeaf5c",
		"df7f00",
		"673b00",
		"472800",
	},
}
