#include <iostream>

using namespace std;

class Node{
    public:
        int value;
        Node* next;

        Node(int value){
            this->value = value;
            this->next = nullptr;
        }
};

class Stack{
    private:
        Node* top;
        Node* bottom;
        int length;
        int capacity;

    public:
        Stack(){
            this->top = nullptr;
            this->bottom = nullptr;
            this->length = 0;
            this->capacity = 0;
        }

        void start(){
            cout << "Please enter the stack capacity: ";
            cin >> this->capacity;
            while(true){
                cout << "Stack management center: " << endl << 
                "choose an option below: " << endl << 
                "1-print top     2-print bottom     3-print length     4-print capcity " << endl << 
                "5-print stack     6-push Node     7-pop a Node" << endl;
                int input;
                cin >> input;
                switch (input){
                    case 1 : this->print_top(); break;
                    case 2 : this->print_bottom(); break;
                    case 3 : this->print_length(); break;
                    case 4 : this->print_capacity(); break;
                    case 5 : this->print_stack(); break;
                    case 6 : this->push_node(); break;
                    case 7 : this->pop_node(); break;
                    default : cout << "No valid option"; break;

                }
                cout << "-----------------------------------------------------" << endl;
            }
        }

        bool is_stack_empty(){
            return (this->top == nullptr || this->length == 0);
        }

        bool is_stack_full(){
            return (this->length == this->capacity);
        }

        void print_top(){
            if (this->is_stack_empty()){
                cout << "There are no Nodes in the stack" << endl;
            }else{
                cout << "Top address: " << this->top << " Top value: " << this->top->value << endl;
            }
        }

        void print_bottom(){
            if (this->is_stack_empty()){
                cout << "THere are no Nodes in the stack" << endl;
            }else{
                cout << "Bottom address: " << this->bottom << " Bottom value: " << this->bottom->value << endl;
            }
        }

        void print_length(){
            cout << "Length: " << this->length << endl;
        }

        void print_capacity(){
            cout << "Capacity " << this->capacity << endl;
        }

        void print_stack(){
            if (this->is_stack_empty()){
                cout << "There are no Nodes in the stack" << endl;
            }else{
                cout << "Stack Elements: " << endl;
                Node* temp = this->bottom;
                while (temp != nullptr){
                    cout << temp->value << " " ;
                    temp = temp->next;
                }
                cout << endl;
            }
        }

        Node* get_node(int index){
            if (this->is_stack_empty()){
                return nullptr;
                cout << "There are no Nodes in the list" << endl;
            }else{
                if (index < 0 || index >= this->length){
                    return nullptr;
                    cout << "Invalid Index" << endl;
                }else{
                    Node* temp = this->bottom;
                    for (int i = 0 ; i < index ; i++){
                        temp = temp->next;
                    }
                    if (temp != nullptr){
                        return temp;
                    }else{
                        cout << "It's null" << endl;
                        return nullptr;
                    }
                }
            }
        }

        void push_node(){
            if (!this->is_stack_full()){
                int value;
                cout << "Please enter the value: " << endl;
                cin >> value;
                Node* new_node = new Node(value);
                if (this->is_stack_empty()){
                    this->bottom = new_node;
                    this->top = new_node;
                }else{
                    this->top->next = new_node;
                    this->top = new_node;
                }
                this->length++;
                cout << "New Node with the value of " << value << " Has been added to the end of the list" << endl;
            }else{
                cout << "The stack is at full capacity" << endl;
            }
        }


        void pop_node(){
            if (this->is_stack_empty()){
                cout << "There are no Nodes in the list" << endl;
            }else{
                Node* temp = this->top;
                Node* previous_node = this->get_node(this->length - 2);
                previous_node->next = nullptr;
                this->top = previous_node;
                this->length--;
                cout << "The last Node with the value of " << temp->value << " Has been removed from the end of the list" << endl;
                delete temp;
                if (this->length == 0){
                    this->bottom = nullptr;
                    this->top = nullptr;
                }
            }
        }
};


int main(){
    Stack* my_stack = new Stack();
    my_stack->start();

}



