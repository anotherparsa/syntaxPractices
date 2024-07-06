void main(){
	Cookie cookie1 = Cookie("Circle", 16.3);
	print(cookie1.is_cooling());
	print(cookie1.shape);
	cookie1.shape = "square";
	print(cookie1.shape);
	Cookie cookie2 = Cookie("Circle", 20.14);
	print(cookie2.shape);
}

class Cookie{
	String? shape;
	double? size;

	Cookie(String shape, double size){
		print("Constructor has been executed");
		this.shape = shape;
		this.size = size;
		this.baking();
	}

	void baking(){
		print("Baking has started");
	}

	bool is_cooling(){
		return false;
	}
}