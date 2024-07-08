void main(){
	Car new_car = Car("V6", 200, 45000);
	new_car.display_top_speed();
	new_car.display_engine_type();
	new_car.display_price();

}

class Car{
	String? engine_type;
	int? top_speed;
	int? price;


	Car(String engine_type, int top_speed, int price){
		print("A car has been made");
		this.engine_type = engine_type;
		this.top_speed = top_speed;
		this.price = price;
	}

	void accelerate(int acceleration){
		print("accelerating ${acceleration} KM/H");
	}

	void display_engine_type(){
		print("engine is ${this.engine_type}");
	}

	void display_price(){
		print("Price is ${this.price}");
	}

	void display_top_speed(){
		print("Top speed is ${this.top_speed}");
	}
}