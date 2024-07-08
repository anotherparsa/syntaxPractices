void main(){
	Car new_car = Car();
	new_car.accelerate(29);
	new_car.display_engine();
}

class Car{
	String engine = "V6";

	void accelerate(int acceleration){
		print("accelerating ${acceleration} KM/H");
	}

	void display_engine(){
		print("engine is ${this.engine}");
	}
}