/// @description Draw probability
// You can write your code in this editor

if move_code == 200 {
	draw_set_color(c_white);
} else {
	draw_set_color(c_red);
}

if message != "" {
	draw_set_font(font_normal);
	draw_text(150, 20, message);	
}
