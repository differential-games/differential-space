/// @description Update scores
// You can write your code in this editor

scoreboard.player_1_score = 0;
scoreboard.player_2_score = 0;
scoreboard.player_3_score = 0;
scoreboard.player_4_score = 0;

var size = ds_list_size(planets);
for (var i = 0; i < size; i++; ) {
	var p = ds_list_find_value(planets, i);
	switch (p.owner) {
	case 1:
		scoreboard.player_1_score++;
		if scoreboard.player_1_score > size * 2 / 3 {
			scoreboard.won = 1;
		}
		break;
	case 2:
		scoreboard.player_2_score++;
		if scoreboard.player_2_score > size * 2 / 3 {
			scoreboard.won = 2;
		}
		break;
	case 3:
		scoreboard.player_3_score++;
		if scoreboard.player_3_score > size * 2 / 3 {
			scoreboard.won = 3;
		}
		break;
	case 4:
		scoreboard.player_4_score++;
		if scoreboard.player_4_score > size * 2 / 3 {
			scoreboard.won = 4;
		}
		break;
	}
}
