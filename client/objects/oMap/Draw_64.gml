/// @description Draw probability
// You can write your code in this editor

if probability >= 0.0 {
	if probability > 80 {
		draw_set_color(c_green);
	} else if probability > 50 {
		draw_set_color(c_yellow);
	} else {
		draw_set_color(c_red);	
	}
	draw_text_transformed(200, 20, "Success: " + string(probability) + "%", 2.0, 2.0, 0.0);
} else if probability = -2 {
	draw_set_color(c_red);	
	draw_text_transformed(200, 20, "Invalid Move", 2.0, 2.0, 0.0)
}
