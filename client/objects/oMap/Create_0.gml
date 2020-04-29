planets = ds_list_create()

for (var i = 0; i < 40; i++) {
	var p = instance_create_depth(0, 0, i, oPlanet);
	p.map_index = i;
	p.parent_map = id;
	ds_list_add(planets, p);
}

post_game = http_post_string("http://localhost:8080/game", "")

turn_counter = instance_create_depth(0, 0, -10, button_next_turn);
turn_counter.parent_map = id;
