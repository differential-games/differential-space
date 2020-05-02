/// @description Update positions
// You can write your code in this editor

var size = ds_list_size(planets);
for (var i = 0; i < size; i++; ) {
	var p = ds_list_find_value(planets, i);
	if (p.a_x == 0 && p.a_y == 0) {
		p.x = 0;
		p.y = 0;
	} else {
		p.x = (origin_x + p.a_x * 32) * scale - 24 + 640;
		p.y = (origin_y + p.a_y * 32) * scale - 24 + 360;	
	}
	
	ds_list_replace(planets, i, p);
}

event_user(2);
