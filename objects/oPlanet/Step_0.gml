/// @description Insert description here
// You can write your code in this editor

if !position_meeting(mouse_x, mouse_y, id) {
	return;
}

if parent_map.selected_planet == -1 {
	return;
}

if parent_map.selected_planet == map_index {
	return;
}

if parent_map.probability >= 0 {
	return;
}

if post_attack_predict != -1 {
	return;
}

parent_map.targeted_planet = map_index;
attack_str = "{\"From\": " + string(parent_map.selected_planet) + ", \"To\": " + string(map_index) + "}"
show_debug_message(attack_str);
post_attack_predict = http_post_string("http://localhost:8080/attack/predict", attack_str);