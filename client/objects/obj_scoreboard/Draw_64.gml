/// @description Insert description here
// You can write your code in this editor

draw_set_alpha(0.80);
draw_set_color(c_black);
draw_rectangle(10, 100, 130, 210, false);
draw_set_alpha(1);

var yDraw = 110;
var yInc = 25;

draw_set_font(font_small);
if player_1_score > 0 {
	draw_set_color(81*256*256 + 166*256 + 0);
	draw_text(20, yDraw, "Player 1: " + string(player_1_score));
	yDraw += yInc;
}
if player_2_score > 0 {
	draw_set_color(36*256*256 + 28*256 + 238);
	draw_text(20, yDraw, "Player 2: " + string(player_2_score));
	yDraw += yInc;
}
if player_3_score > 0 {
	draw_set_color(166*256*256 + 84*256 + 0);
	draw_text(20, yDraw, "Player 3: " + string(player_3_score));
	yDraw += yInc;
}
if player_4_score > 0 {
	draw_set_color(145*256*256 + 45*256 + 102);
	draw_text(20, yDraw, "Player 4: " + string(player_4_score));
	yDraw += yInc;
}
if won != -1 {
	draw_set_color(c_white);
	draw_text(20, yDraw, "Player " + string(won) + " wins!");
}
