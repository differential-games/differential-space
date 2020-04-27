/// @description Insert description here
// You can write your code in this editor
if ds_map_find_value(async_load, "id") == post_attack {
	with(parent_map){ event_user(0); }
	return;
} else if ds_map_find_value(async_load, "id") == post_attack_predict {
	var status = ds_map_find_value(async_load, "http_status");
	show_debug_message("Status: " + string(status));
	if status != 200 {
		parent_map.probability = -2;
	} else {
		var r_str = ds_map_find_value(async_load, "result");
		var resultMap = json_decode(r_str);
		parent_map.probability = 100 * ds_map_find_value(resultMap, "P");	
	}
	post_attack_predict = -1;
} else {
	return;
}
