/// @description Attack
// You can write your code in this editor

if parent_map.selected_planet != -1 and parent_map.selected_planet != map_index {
	attack_str = "{\"From\": " + string(parent_map.selected_planet) + ", \"To\": " + string(map_index) + "}"
	post_attack = http_post_string("http://localhost:8080/attack", attack_str);
	
	parent_map.selected_planet = -1;
}
