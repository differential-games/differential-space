/// @description Insert description here
// You can write your code in this editor

draw_set_alpha(0.80);
draw_set_color(c_black);
draw_rectangle(10, 100, 130, 210, false);
draw_set_alpha(1);

var yDraw = 110;
var yInc = 25;

var size = ds_list_size(player_scores);

draw_set_color(c_white)
for (var i = 1; i <= size; i++; ) {
	var player_score = ds_list_find_value(player_scores, i-1);
	if (player_score == 0) {
		continue;
	}
	if parent_map.turn_counter.turn == i {
		draw_set_font(font_small_bold);
		draw_text(20, yDraw, "* Player "+ string(i) + ": " + string(player_score));
	} else {
		draw_set_font(font_small);
		draw_text(20, yDraw, "Player "+ string(i) + ": " + string(player_score));
	}
	yDraw += yInc;	
}

if won != -1 {
	draw_set_color(c_white);
	draw_text(20, yDraw, "Player " + string(won) + " wins!");
}
