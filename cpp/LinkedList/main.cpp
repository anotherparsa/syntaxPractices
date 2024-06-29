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

        void start(){
            while(true){
                cout << "Linked List management center: " << endl <<
                "choose an option below:" << endl <<
                "0-check if list is empty?     1-print head     2-print tail     3-print length     4-print list " << endl <<
                "5-print specific Node value by index     6-set specific Node value     7-insert Node at a specific index     8-delete Node at a specific index " << endl << 
                "9-append Node     10-delete last Node     11-prepend Node     12-delete first Node " << endl << 
                "13-purge List     14-check if list is sorted" << endl;
                int input;
                cin >> input;
                switch (input){
                    case 0 : cout << is_list_empty() << endl; break;
                    case 1 : this->print_head(); break;
                    case 2 : this->print_tail(); break;
                    case 3 : this->print_length(); break;
                    case 4 : this->print_list(); break;
                    case 5 : this->print_node_value(); break;
                    case 6 : this->set_node_value(); break;
                    case 7 : this->insert_node(); break;
                    case 8 : this->delete_node(); break;            
                    case 9 : this->append_node(); break;
                    case 10 : this->delete_last_node(); break;
                    case 11 : this->prepend_node(); break;
                    case 12 : this->delete_first_node(); break;
                    case 13 : this->purge_list(); break;
                    case 14 : this->is_list_sorted(); break;
                    default : cout << "No valid option"; break;
                }
                cout << "-----------------------------------------------------" << endl;
            }
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
            }else{
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
                cout << "There are no Nodes in the list" << endl;
            }else{
                if (index < 0 || index >= this->length){
                    return nullptr;
                    cout << "Invalid Index" << endl;
                }else{
                    Node* temp = this->head;
                    for (int i = 0 ; i < index ; i++){
                        temp = temp->next;
                    }
                    return temp;
                }
            }
        }

        void append_node(){
            int value;
            cout << "Please enter the value: ";
            cin >> value;
            Node* new_node = new Node(value);
            if (this->is_list_empty()){
                this->head = new_node;
                this->tail = new_node;
            }else{
                this->tail->next = new_node;
                this->tail = new_node;
            }
            this->length++;
            cout << "New Node with the value of " << value << " Has been added to the end of the list" << endl;
        }

        void prepend_node(){
            int value;
            cout << "Please enter the value: ";
            cin >> value;
            Node* new_node = new Node(value);
            if (this->is_list_empty()){
                this->head = new_node;
                this->tail = new_node;
            }else{
                new_node->next = this->head;
                this->head = new_node;
            }
            this->length++;
            cout << "New Node with the value of " << value << " Has been added to the beginning of the list" << endl;
        }

        void delete_first_node(){
            if (this->is_list_empty()){
                cout << "There are no Nodes in the list" << endl;
            }else{
                Node* temp = this->head;
                this->head = this->head->next;
                cout << "Node with the value of " << temp->value << " Has been removed from the beginning of the list" << endl;
                delete temp;
                this->length--;
                if (this->length == 0){
                    this->head = nullptr;
                    this->tail = nullptr;
                }
            }
        }

        void delete_last_node(){
            if (this->is_list_empty()){
                cout << "There are no Nodes in the list" << endl;
            }else{
                Node* temp = this->tail;
                Node* previous_node = this->get_node(this->length - 2);
                previous_node->next = nullptr;
                this->tail = previous_node;
                this->length--;
                cout << "The last Node with the value of " << temp->value << " Has been removed from the end of the list" << endl;
                delete temp;
                if (this->length == 0){
                    this->head = nullptr;
                    this->tail = nullptr;
                }
            }
        }

        void print_node_value(){
            int index;
            cout << "Please enter the index: ";
            cin >> index;
            if (this->is_list_empty()){
                cout << "There are no Nodes in the list" << endl;
            }else{
                if (index < 0 || index >= this->length){
                    cout << "Invalid Index" << endl;
                }else{
                    cout << "The Node at the index of " << index << " Has the value of " << this->get_node(index)->value << endl;
                }
            }
        }

        void set_node_value(){
            int value;
            int index;
            cout << "Please enter the index: ";
            cin >> index;
            cout << "Please enter the value: ";
            cin >> value;
            if (this->is_list_empty()){
                cout << "There are no Nodes in the list" << endl;
            }else{
                if (index < 0 || index >= this->length){
                    cout << "Invalid Index" << endl; 
                }else{
                    Node* temp = this->get_node(index);
                    cout << "The Node at index " << index << " Had the value of " << temp->value ;
                    temp->value = value;
                    cout << " Now it has the value of " << value << endl;
                }
            }
        }

        void insert_node(){
            int value;
            int index;
            cout << "Please enter the index: ";
            cin >> index;           
            cout << "Please enter the value: ";
            cin >> value;
            if (index < 0 || index < this->length){
                cout << "Invalid Index " << endl;
            }else{
                if (index == 0 ){
                    this->prepend_node();
                }else if (index == this->length){
                    this->append_node();
                }else{
                    Node* new_node = new Node(value);
                    Node* previous_node = this->get_node(index - 1);
                    new_node->next = previous_node->next;
                    previous_node->next = new_node;
                    this->length++;
                    cout << "New Node with the value of " << value << " Has been added at index " << index << endl;
                }
            }
        }

        void delete_node(){
            int index;
            cout << "Please enter the index ";
            cin >> index;
            if (this->is_list_empty()){
                cout << "There are no Nodes in the list" << endl;
            }else{
                if (index < 0 || index >= this->length){
                    cout << "Invalid Index " << endl;
                }else{
                    if (index == 0){
                        this->delete_first_node();
                    }else if (index == this->length - 1){
                        this->delete_last_node();
                    }else{
                        Node* temp = this->get_node(index);
                        Node* previous_node = this->get_node(index - 1);
                        previous_node->next = temp->next;
                        this->length--;
                        cout << "Node with the value of " << temp->value << " Has been removed at index " << index << endl;
                        delete temp;
                    }
                }
            }
        }

        void purge_list(){
            if (this->is_list_empty()){
                cout << "There are no Nodes in the list" << endl;
            }else{
                cout << "Are you sure to purge the list? enter \"Yes\"" << endl;
                string Answer;
                cin >> Answer;
                if (Answer == "Yes"){
                    cout << "Purging List: " << endl; 
                    while (this->length != 0){
                        this->delete_first_node();
                        this->length--;
                    }
                }else{
                    cout << "OK" << endl;
                }
            }
        }

        void is_list_sorted(){
            if (this->is_list_empty()){
                cout << "There are no Nodes in the list" << endl;
            }else{
                Node* first = this->head;
                Node* second = this->head->next;
                int flag = 1;
                while (second != nullptr){
                    if (first->value > second->value){
                        flag = 0;
                        break;
                    }
                    first = second;
                    second = second->next;
                }
                if (flag == 1){
                    cout << "The list is sorted " << endl;
                }else{
                    cout << "The list is not sorted " << endl;
                }
            }
        }

};

int main(){
    LinkedList* my_linked_list = new LinkedList();
    my_linked_list->start();

}
