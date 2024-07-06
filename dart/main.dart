void main(){
	Cookie cookie1 = Cookie();
	cookie1.baking();
	print(cookie1.is_cooling());
	print(cookie1.shape);
	cookie1.shape = "square";
	print(cookie1.shape);
}

class Cookie{
	String shape = "circle";
	double size = 15.2;

	void baking(){
		print("Baking has started");
	}

	bool is_cooling(){
		return false;
	}
}