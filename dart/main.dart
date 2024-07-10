void main(){
	Vehicle new_vehicle = Vehicle();
	new_vehicle.accelerate(20);
	new_vehicle.get_speed();
	new_vehicle.accelerate(30);
	new_vehicle.get_speed();
	new_vehicle.deccelerate(19);
	new_vehicle.get_speed();

	Car new_car = Car();
	new_car.accelerate(50);
	new_car.get_speed();
	new_car.number_of_wheel = 5;
	print(new_car.number_of_wheel);
	new_car.number_of_wheel = 10;
	print(new_car.number_of_wheel);


}

class Vehicle{
	int speed = 0;
	bool is_running = false;

	void accelerate(int acceleration){
		this.speed += acceleration;
		print("Vehicle is accelerating");
		this.is_running = true;
	}

	void get_speed(){
		print("Vehicle's speed is ${this.speed}");
	}

	void deccelerate(int decceleration){
		if (this.speed - decceleration < 0){
			print("You can't decelerate that much");
		}else{
			this.speed -= decceleration;
			print("Vehicle is deccelerating");
			if (this.speed == 0){
				this.is_running = false;
			}
		}
	}

	void check_is_running(){
		if (this.is_running){
			print("It's is running");
		}else{
			print("It's not running");
		}
	}
}

class Car extends Vehicle{
	int? number_of_wheel;
}