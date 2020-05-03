/// @description Insert description here
// You can write your code in this editor
if ds_map_find_value(async_load, "id") == post_button {
	if ds_map_find_value(async_load, "status") != 0	{
		// There was some sort of server error?
		return;
	}
	var turn_str = ds_map_find_value(async_load, "result");
	turn = int64(turn_str);
	
	with(parent_map){ event_user(0); }
	return;
} else {
	return;
}