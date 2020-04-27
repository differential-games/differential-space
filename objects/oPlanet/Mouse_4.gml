/// @description Select/deselect
// You can write your code in this editor

if parent_map.selected_planet == map_index {
	parent_map.selected_planet = -1;
} else {
	parent_map.selected_planet = map_index;
}
