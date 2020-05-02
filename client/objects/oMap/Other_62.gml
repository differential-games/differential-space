/// @description Handle get map
// You can write your code in this editor

var async_id = ds_map_find_value(async_load, "id");

if async_id == get_planets {
	if ds_map_find_value(async_load, "status") == 0	{
		tries = 0;
		var r_str = ds_map_find_value(async_load, "result");
	} else {
		tries = tries + 1;
		if tries > 5 {
			show_error("unable to contact server", true);
		}
		event_user(0);
		return;
	}
} else if async_id == post_next || async_id == post_game {
	event_user(0);
	return;
} else {
	return;
}

var resultMap = json_decode(r_str);
var new_planets = ds_map_find_value(resultMap, "default")

var size = ds_list_size(planets);
var newSize := ds_list_size(new_planets);

for (var i = 0; i < size; i++; ) {
	var op = ds_list_find_value(planets, i);
	if (i >= newSize) {
		op.a_x = 0;
		op.a_y = 0;
		ds_list_replace(planets, i, op);
		continue;
	}
	
	var p = ds_list_find_value(new_planets, i);
	// Set absolute x and y values.
	var x_str = ds_map_find_value(p, "X");
	op.a_x = real(x_str);

	var y_str = ds_map_find_value(p, "Y")
	op.a_y = real(y_str);
	
	var owner = ds_map_find_value(p, "Owner");
	op.owner = real(owner);
	
	var ready = ds_map_find_value(p, "Ready");
	op.ready = bool(ready);
	
	var strength = ds_map_find_value(p, "Strength");
	op.strength = real(strength);
	
	ds_list_replace(planets, i, op);
}
ds_list_destroy(new_planets)

event_user(1);
