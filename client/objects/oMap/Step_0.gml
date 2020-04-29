/// @description Insert description here
// You can write your code in this editor

if mouse_check_button_pressed(mb_middle) {
	dragging = true;
	origin_drag_x = origin_x;
	origin_drag_y = origin_y;
	mouse_drag_x = mouse_x;
	mouse_drag_y = mouse_y;
}

if mouse_check_button_released(mb_middle) {
	dragging = false;
}

if dragging {
	origin_x = origin_drag_x + (mouse_x - mouse_drag_x) / scale;
	origin_y = origin_drag_y + (mouse_y - mouse_drag_y) / scale;
	event_user(1);
}

if message != "" {
	if targeted_planet == -1 {
		message = "";
	} else if !position_meeting(mouse_x, mouse_y, ds_list_find_value(planets, targeted_planet)) {
		message = "";
	}
}
