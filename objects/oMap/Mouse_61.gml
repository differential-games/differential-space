/// @description Zoom out.
// You can write your code in this editor
if scale * 0.95 <= 0.25 {
	return;
}
scale = max(0.25, scale * 0.95);

origin_x = origin_x + (mouse_x - 960) * (1/0.95 - 1) / scale;
origin_y = origin_y + (mouse_y - 540) * (1/0.95 - 1) / scale;

event_user(1);
