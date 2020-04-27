/// @description Insert description here
// You can write your code in this editor

draw_sprite(spr_planet, owner, x, y);

if parent_map.selected_planet == map_index {
	draw_sprite(spr_select, 0, x, y);
} else if parent_map.targeted_planet == map_index && parent_map.probability >= 0 {
	draw_sprite(spr_select, 0, x, y);
} else if parent_map.selected_planet == -1 && position_meeting(mouse_x, mouse_y, id) {
	draw_sprite(spr_select, 0, x, y);
}

if parent_map.selected_planet == map_index || position_meeting(mouse_x, mouse_y, id) {
	draw_set_color(c_gray);
	draw_set_circle_precision(64);
	draw_circle(x+20, y+20, 5 * 32 * parent_map.scale, true);
}

if !ready {
	draw_sprite(spr_done, 0, x, y)
}

if strength > 0 {
	draw_sprite(spr_strength, strength-1, x, y);
}
