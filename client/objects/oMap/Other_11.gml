/// @description Update positions
// You can write your code in this editor

var size = ds_list_size(planets);
for (var i = 0; i < size; i++; ) {
	var p = ds_list_find_value(planets, i);
	p.x = (origin_x + p.a_x * 32) * scale - 24 + 600;
	p.y = (origin_y + p.a_y * 32) * scale - 24 + 450;
	
	ds_list_replace(planets, i, p);
}
