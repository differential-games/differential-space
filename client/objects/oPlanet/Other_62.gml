/// @description Insert description here
// You can write your code in this editor
if ds_map_find_value(async_load, "id") == post_attack {
	with(parent_map){ event_user(0); }
	return;
} else if ds_map_find_value(async_load, "id") == post_attack_predict {
	parent_map.message = ds_map_find_value(async_load, "result");
	parent_map.move_code = ds_map_find_value(async_load, "http_status");
	post_attack_predict = -1;
} else {
	return;
}
