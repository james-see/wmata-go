import 'regenerator-runtime/runtime';
import axios from 'axios';

const BASE_URL = 'https://api.wmata.com/StationPrediction.svc/json/GetPrediction/C05';

const getTodoItems = async () => {
  try {
    const response = await axios.get(`${BASE_URL}`, {
        headers: {
            'api_key': '2522cab0b36a45bcb592b8e621028ed0'
        }
    });

    const todoItems = response.data;

    console.log(`GET: Here's the list of todos`, todoItems.Trains);

    return todoItems.Trains;
  } catch (errors) {
    console.error(errors);
  }
};

const createTodoElement = item => {
    const todoElement = document.createElement('li');
    todoElement.appendChild(document.createTextNode(item));
  
    return todoElement;
  };

  const updateTodoList = todoItems => {
    const todoList = document.querySelector('ul');
  
    if (Array.isArray(todoItems.Trains) && todoItems.Trains.length > 0) {
      todoItems.Trains.map(todoItem => {
        let jsondata = todoItem.json()
        todoList.appendChild(createTodoElement(jsondata.Car));
      });
    } else if (todoItems) {
        todoItems.forEach(car => {
            if (car.DestinationCode == "K08" || car.DestinationCode == "N12") {
            todoList.appendChild(createTodoElement("Destination: " + car.Destination + " Status: " + car.Min));
            }
        }
        );
    }
  };
  
  const main = async () => {
    updateTodoList(await getTodoItems());
  };
  
  main();


