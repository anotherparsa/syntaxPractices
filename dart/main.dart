void main(){
	Cookie cookie1 = Cookie();
	cookie1.baking();
	print(cookie1.is_cooling());
	print(cookie1.shape);
	cookie1.shape = "square";
	print(cookie1.shape);
	Cookie cookie2 = Cookie();
}

class Cookie{
	String shape = "circle";
	double size = 15.2;

	Cookie(){
		print("Constructor has been executed");
	}

	void baking(){
		print("Baking has started");
	}

	bool is_cooling(){
		return false;
	}
}