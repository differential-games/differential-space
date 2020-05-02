planets = ds_list_create()

for (var i = 0; i < 96; i++) {
	var p = instance_create_depth(0, 0, i, oPlanet);
	p.map_index = i;
	p.parent_map = id;
	ds_list_add(planets, p);
}

turn_counter = instance_create_depth(0, 0, -10, button_next_turn);
turn_counter.parent_map = id;

scoreboard = instance_create_depth(0, 0, -10, obj_scoreboard);
scoreboard.parent_map = id;

post_game = http_post_string("http://localhost:8080/game", "")
