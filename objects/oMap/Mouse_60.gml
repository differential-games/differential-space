/// @description Zoom in.
// You can write your code in this editor
if scale * 1.05 >= 4.00 {
	return;
}
scale = min(4.00, scale * 1.05);

origin_x = origin_x + (mouse_x - 960) * (1/1.05 - 1) / scale;
origin_y = origin_y + (mouse_y - 540) * (1/1.05 - 1) / scale;

event_user(1);
