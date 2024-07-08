void main(){
	Car new_car = Car();
	new_car.accelerate(20);
	new_car.get_speed();
	new_car.accelerate(30);
	new_car.get_speed();
	new_car.deccelerate(19);
	new_car.get_speed();

}

class Car{
	int speed = 0;
	bool is_running = false;

	void accelerate(int acceleration){
		this.speed += acceleration;
		print("Car is accelerating");
		this.is_running = true;
	}

	void get_speed(){
		print("Car's speed is ${this.speed}");
	}

	void deccelerate(int decceleration){
		if (this.speed - decceleration < 0){
			print("You can't decelerate that much");
		}else{
			this.speed -= decceleration;
			print("Car is deccelerating");
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