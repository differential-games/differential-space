/// @description Update scores
// You can write your code in this editor

var size = ds_list_size(scoreboard.player_scores);
for (var i = 0; i < size; i++; ) {
	ds_list_set(scoreboard.player_scores, i, 0);
}

var size = ds_list_size(planets);
for (var i = 0; i < size; i++; ) {
	var p = ds_list_find_value(planets, i);
	if (p.owner == 0) {
		continue;
	}
	var newScore = ds_list_find_value(scoreboard.player_scores, p.owner-1) + 1;
	if newScore > (size * 2 / 3) {
		scoreboard.won = p.owner;
	}
	ds_list_set(scoreboard.player_scores, p.owner-1, newScore);
}
