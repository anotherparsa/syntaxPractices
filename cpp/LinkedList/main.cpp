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

class LinkedList{
    private:
        Node* head;
        Node* tail;
        int length;

    public:
        LinkedList(){
            this->head = nullptr;
            this->tail = nullptr;
            this->length = 0;
        }

        bool is_list_empty(){
            return (this->head == nullptr || this->length == 0);
        }

        void print_head(){
            if (this->is_list_empty()){
                cout << "There are no Nodes in the list" << endl;
            }else{
                cout << "Head address: " << this->head << " Head value: " << this->head->value << endl;
            }
        }

        void print_tail(){
            if (this->is_list_empty()){
                cout << "There are no Nodes in the list" << endl;
            }eles{
                cout << "Tail address: " << this->tail << "Tail value: " << this->tail->value << endl;
            }
        }

        void print_length(){
            cout << "Length: " << this->length << endl;
        }

        void print_list(){
            if (this->is_list_empty()){
                cout << "There are no Nodes in the list" << endl;
            }else{
                cout << "List Elements: " << endl;
                Node* temp = this->head;
                while(temp != nullptr){
                    cout << temp->value << " " ;
                    temp = temp->next;
                }
                cout << endl;
            }
        }

        Node* get_node(int index){
            if (this->is_list_empty()){
                return nullptr;
            }else{
                if (index < 0 || index >= this->length){
                    return nullptr;
                }else{
                    Node* temp = this->head;
                    for (int i = 0 ; i < index ; i++){
                        temp = temp->next;
                    }
                    return temp;
                }
            }
        }
};

int main(){
    LinkedList* my_linked_list = new LinkedList();
    cout << my_linked_list->is_list_empty() << endl;
}
