/// @description Insert description here
// You can write your code in this editor

var idx = turn * 2;

if position_meeting(mouse_x, mouse_y, id) {
	draw_sprite(spr_next_turn, idx + 1, 10, 10);
} else {
	draw_sprite(spr_next_turn, idx, 10, 10);
}

if turn == 0 {
	draw_set_color(c_white);
	draw_set_font(font_large)
	draw_text(x+27, y+14, "Start");
}
