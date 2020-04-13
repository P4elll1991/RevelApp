class bookTab {

    constructor(){
    this.modal = new modalBook();
    this.modal.giveData(this.modal);
    
  }
  
  // кнопки управления
    buttons = [
      { id:"change", view:"button", type:"icon", icon:"mdi mdi-pen", value: "Изменить"},                           
      { id:"push",  view:"button", type:"icon", icon:"mdi mdi-plus-box-outline", value: "Добавить"},
      { id:"goToEmployee", view:"button", type:"icon", icon:"mdi mdi-account", value: "Перейти к сотруднику"},
      { id:"delete", view:"button", type:"icon", icon:"mdi mdi-delete-forever", value: "Удалить"},      
      { id:"GiveOut", view:"button", value: "Выдать"},      
      { id:"Return", view:"button", value: "Вернуть"},      
      ];
  

      // столбцы таблицы
      columns = [
        { id:"ch1", header:{ content:"masterCheckbox", contentId:"mc1" }, template:"{common.checkbox()}", width: 50,},
        { id:"Isbn",    header:"ISBN", adjust:true, sort: "int",},
        { id:"BookName",   header:"Название", adjust:true, sort: "string",},
        { id:"Autor",  header:"Автор", adjust:true, sort: "string"},
        { id:"Publisher",  header:"Издатель", adjust: true, sort: "string"},
        { id:"Year",  header:"Год", adjust: true, sort: "int",},
        { id:"Status",  header:"Статус", width:150, sort: "string"},
        { id:"Name",  header:"Сотрудник", width:200, sort: "string"},
        { id:"Cellnumber",  header:"Номер телефона", width:200, sort: "string"},
        { id:"Datestart",  header:"Дата выдачи", adjust: true, format:webix.Date.dateToStr("%d.%m.%Y"), sort: "date"},
        { id:"Datefinish",  header:"Дата сдачи", adjust: true, format:webix.Date.dateToStr("%d.%m.%Y"), sort: "date"}
                    ];
  
    init() { // инициализация таблицы
    
      this.view = {
        view:"layout",
        padding:10,
        id: "bookView", 
        type: "wide",
            rows:[
            

            // меню
  
            { id:"bookSidebar", select:false,
              cols: this.buttons},
            
            // Таблица
  
            {
            view:"datatable", 
            id:"bookTable", 
            wordBreak: "break-all", 
            css:"webix_data_border webix_header_border",
            multiselect:true, 
            columns: this.columns, 
            select:true, 
            },

  
          ],
                   
      }
      return this.view;
  
    }

    initWindow() { // инициализация модального окна
      this.up = new windowBook();
      this.window = this.up.getWindow();
      return this.window;
    }

    getView() {
      return this.init(); // отправка таблицы
    }

    // прикрепление событий

    editeEvents(parent){ 
      var options = new windowBook(); // отправка опций в форму
      options.optionsBook();
      // удаление
      $$("delete").attachEvent("onItemClick", function(){
        parent.delete();
      });

      // открытие окна изменений

      $$("change").attachEvent("onItemClick", function(){
        var item = $$("bookTable").getSelectedItem();
       
        var item_data = $$("formBook").getValues()
        console.log(item_data);
        var check = item_data.Status;
        console.log(check);
        if(check == "Нет в наличии") {
          webix.confirm({
            text: "Нельзя редактировать книгу пока она не будет сдана", 
            ok: "OK",
          }).then(function(){
            return
          });
          return;
        };
        $$("Status").hide();
        $$("Name").hide();
        parent.checkWin = true; // флаг который сигнализирует модальное окно
        var item_data = $$("formBook").getValues()
        var check = item_data.BookName;
        console.log(check);
        parent.formBlock(true, false);
        if(check != "")
             $$("windowBook").show();
      });

      // открытие окна изменений по двойному щекчку по эементу

      $$("bookTable").attachEvent("onItemDblClick", function(){
        $$("Status").hide();
        parent.checkWin = true; // флаг который сигнализирует модальное окно
        var item_data = $$("formBook").getValues()
        var check = item_data.name;
        parent.formBlock(true, false);
        if(check != "")
             $$("windowBook").show();
      });

      // открытие окна добавления

      $$("push").attachEvent("onItemClick", function(){
        parent.checkWin = false; // флаг который сигнализирует модальное окно
        $$("formBook").clear();
        $$("Status").hide();
        parent.formBlock(false, false);
        $$("windowBook").show();
      });

      // выдать элемент

      $$("GiveOut").attachEvent("onItemClick", function(){
        var item = $$("bookTable").getSelectedItem();
        
        var item_data = $$("formBook").getValues()
        console.log(item_data);
        var check = item_data.Status;
        console.log(check);
        if(check != "В наличии") return; // если нет в наличии то отдать нельзя
        $$("Status").setValue("Нет в наличии"); // смена статуса на нет в наличии в форме
        $$("Status").hide();
        $$("windowBook").show();
        parent.checkWin = true;
        var item_data = $$("formBook").getValues()
        var check = item_data.BookName;
        parent.formBlock(true, true);
        console.log(check);
        if(check != "")
             $$("windowBook").show();
      });

      // операции после выбора элемента
      $$("bookTable").attachEvent("onAfterSelect", function(){
        parent.afterSelect();
     });

     // оперции после снятия выбора

      $$("bookTable").attachEvent("onAfterUnSelect", function(selection){

           parent.afterUnSelect(selection);
      });

      // вернуть книгу

      $$("Return").attachEvent("onItemClick", function(){
        var item = $$("bookTable").getSelectedItem();
        
        parent.checkWin = true;
        $$("Status").setValue("В наличии");
        $$("formBook").elements["Status"].refresh();
        var form = $$("formBook");
        var item_data = form.getValues();
        console.log(item_data);
        parent.updateTab(parent.checkWin, parent);
           
      });

      //выйти из окна

      $$("exitWindowBook").attachEvent("onItemClick", function() {
          $$("windowBook").hide();
          $$("formBook").clear();
          $$("formBook").clearValidation();
          $$("Status").show();

          parent.afterSelect();
      });

      // подтвердить отправку формы

      $$("updateBookTab").attachEvent("onItemClick", function(){
          
          parent.updateTab(parent.checkWin, parent);

      });

      // событие после смены статуса в форме

      $$("formBook").elements["Status"].attachEvent("onChange", function(newv, oldv){
        if (newv == "Нет в наличии") {
           $$("formBook").elements["Name"].show();


        } else if (newv == "В наличии"){
          $$("formBook").elements["Name"].hide();
        }
        });

        // перейти к сотруднку
      $$("goToEmployee").attachEvent("onItemClick", function(){
        parent.focus();
      });

      $$("bookTable").attachEvent("onCheck", function(rowId, colId, state){
        if (state == 1) {
          $$("bookTable").select(rowId, true);
        } else if (state == 0){
          $$("bookTable").unselect(rowId);
        }
          
      });
  
    }

    // удалить книгу

    delete(){
    var list = $$("bookTable");
    var item_id = list.getSelectedId();
    var item = list.getSelectedItem();
    console.log(item.Status);
    var IdList = []; 
    // если выбрана одна книга
    if (!Array.isArray(item)) {
      if (item_id){
        if (item.Status == "Нет в наличии") { // проверка не находится ли книга у сотрудника
          webix.confirm({
            text: "Нельзя удалить книгу пока она не будет сдана", 
            ok: "OK",
          }).then(function(){
            return
          });
        } else { // удаление
          webix.confirm({
            text: "Вы действительно хотите удалить книгу?",
            cancel: "Нет", 
            ok: "Да",
          }).then(function(){
            list.remove(item_id);
            IdList.push(item.Id);
            console.log(IdList);
            webix.ajax().headers({
              "Content-type":"application/json"
              }).post("/Books/Delete", JSON.stringify(IdList)).then(function(data){
              data = data.json();
              data.forEach(function(val){
                let obj = new modalBook()
                val = obj.dataProcessing(val);
              });
              $$("bookTable").parse(data);
              });
          });
        } 
      }
    } 
    else { // если выбрано много книг
      var i = 0; 
      item.forEach(function(val){ // проверка не находится ли книга у сотрудника создание массива id
        i++; // счетчик
        if (val.Status == "Нет в наличии") {
          webix.confirm({
            text: "Нельзя удалить книгу пока она не будет сдана.", 
            ok: "OK",
          }).then(function(){
            return
          });
        } else {
          IdList.push(val.Id);
        }
        
      });
      if (item_id && (IdList.length == i)){ // если длина массива совпадает со значение счетчика есть id то производится удаление
        webix.confirm({
            text: "Вы действительно хотите удалить книги?",
            cancel: "Нет", 
            ok: "Да",
          }).then(function(){
            IdList.forEach(function(val){
              list.remove(val);
            });
            console.log(IdList);
            webix.ajax().headers({
              "Content-type":"application/json"
              }).post("/Books/Delete", JSON.stringify(IdList)).then(function(data){
             data = data.json();
             data.forEach(function(val){
              let obj = new modalBook()
              val = obj.dataProcessing(val);
            });
            $$("bookTable").parse(data);
            });
          });
      }
    }
  }

  // событие после выбора элемента

  afterSelect() {

      var item = $$("bookTable").getSelectedItem();
      let check = false;
      for (let key in item) {
        if (key == "Id") {
          check = true;
        }
      }
      if(check == true){
        for (let key in item) {
          if (key == "Id") {
            break;
          }
          delete item[key];
        }
      }
      
      console.log(item);
      if (!item) return;
      var x = item.Name;
      item.Name = item.Employeeid; // Нужно чтобы в форме выбиралась нужная опция
      $$("formBook").setValues(item); // автозаполнение формы
      $$("formBook").setValues(item);
      if (Array.isArray(item)) { // проверка один или несколько елементов выбрано
        item.forEach(function(val){ // если несколько
          val.ch1 = 1; // нажатие чекбокса
          item.Name = x;
          $$("bookTable").updateItem(val.id, val);
        });
        return;
      }
      item.ch1 = 1;
      item.Name = x;
      $$("bookTable").updateItem(item.id, item);
    }

// функция после снятие выделения
    afterUnSelect(selection){
      var item = selection;
      console.log(item);
      item.ch1 = 0; // снятие чекбокса
      if(!item.id) return;
      $$("bookTable").updateItem(item.id, item);
    }



// Изменение и добавление данных    
  updateTab(check, parent){ // check - флаг указывающий что будет происходитьЖ изменение или удаление

// Получение данных формы
    var table = $$("bookTable");
    var item = table.getSelectedItem();
     var form = $$("formBook");
     var item_data = form.getValues();
     
    
// Проверка валидности данных
    if (!form.validate()){
        webix.message({ type:"error", text:"Некорректно заполненная форма" });
        return
    }
     if(!check) { // если добавляем книгу
      parent.addBook(item_data); // добавить книгу
     } else { // иначе изменяем данные
      parent.updateBook(item, item_data, parent);
     }
     
     $$("windowBook").hide();
     form.clear();
  }
// перевод фокуса к сотруднику
  focus() {
    var item = $$("bookTable").getSelectedItem();
    if (!item) return;
    var item_id = item.id;
    var focusId = item.Employeeid;
  
    if (!focusId) return;

    $$("bookTable").unselect(item_id);
    item.ch1 = 0;
    $$("bookTable").updateItem(item.id, item);
    $$("staffTable").unselectAll();
    $$("staffTable").select(focusId,true);
    $$("staffView").show();
    $$("staffTable").showItem(focusId);
  }
// блокировка и снятие блокировки с полей формы в зависимости от того какое окно выбрано
  formBlock(bool1, bool2){
    $$("formBook").elements["Isbn"].config.readonly = bool1;
        $$("formBook").elements["Isbn"].refresh();
        $$("formBook").elements["BookName"].config.readonly = bool2;
        $$("formBook").elements["BookName"].refresh();
        $$("formBook").elements["Autor"].config.readonly = bool2;
        $$("formBook").elements["Autor"].refresh();
        $$("formBook").elements["Publisher"].config.readonly = bool2;
        $$("formBook").elements["Publisher"].refresh();
        $$("formBook").elements["Year"].config.readonly = bool2;
        $$("formBook").elements["Year"].refresh();
        $$("formBook").elements["Status"].config.readonly = bool2;
        $$("formBook").elements["Status"].refresh();
  }

  //функция добавления книги

  addBook(item_data){
    if (item_data.id) {
      for (var i in this.books){
      if(item_data.id == this.books[i].id) {
        webix.message({ type:"error", text:"Книга с таким ISBN уже существует" });
        return
      };
    };
  }

  item_data.Isbn = Number(item_data.Isbn);
  item_data.Year = Number(item_data.Year);

  webix.ajax().headers({
    "Content-type":"application/json"
}).post("/Books/Add", JSON.stringify(item_data)).then(function(data){
  data = data.json();
  data.forEach(function(val){
    let obj = new modalBook()
    val = obj.dataProcessing(val);
  });
  $$("bookTable").parse(data);
  });
  }

  // функция добавления события

  addEvent(item, item_data) {
    $$("staffTable").eachRow(function(row){ // формируем данные для создания события
      var record = $$("staffTable").getItem(row);
      if (record.Id == item_data.Name){
        item_data.Cellnumber = record.Cellnumber;
        item_data.Name = record.Name;
      }
    });
    this.postDataEvent = {
      Event: "Возвращено",
      BookId :Number(item_data.Id),
      BookNameJ: item_data.BookName,
      IsbnJ: Number(item_data.Isbn),
    };

    if(item_data.Status == "В наличии"){ // проверяем какое было событие и составляем данные исходя из этого
      console.log("Возвращено")
      this.postDataEvent.Event = "Возвращено";
      this.postDataEvent.NameJ = item.Name;
      this.postDataEvent.CellnumberJ = Number(item.Cellnumber);
      this.postDataEvent.EmployeeId = Number(item.Employeeid);
      console.log(this.postDataEvent)
    } else {
      console.log("Выдано")
      this.postDataEvent.Event = "Выдано";
      this.postDataEvent.NameJ = item_data.Name;
      this.postDataEvent.CellnumberJ = Number(item_data.Cellnumber);
      this.postDataEvent.EmployeeId = Number(item_data.Name);
      console.log(this.postDataEvent)
    }

    webix.ajax().headers({ // запрос на создание события в журнале
      "Content-type":"application/json"
    }).post("/Journal/Add", JSON.stringify(this.postDataEvent));
  }

// Обновление данных книги

  updateBook(item, item_data, parent) {
    item_data.Id = Number(item_data.Id); // конверция текстовых данных таблицы
    item_data.Isbn = Number(item_data.Isbn);  
    item_data.Year = Number(item_data.Year);
    if (item_data.Status == "В наличии") {
        item_data.EmployeeId = 1 

    } else {
      item_data.EmployeeId = Number(item_data.Name); 
    }
  
    if (item.Status != item_data.Status){ // проверям было ли событие выдачи/возвращения
      parent.addEvent(item, item_data);
    } 
    
      webix.ajax().headers({
        "Content-type":"application/json"
    }).post("/Books/Update", JSON.stringify(item_data)).then(function(data){
      data = data.json();
      console.log(data);
      data.Books.forEach(function(val){
          let obj = new modalBook()
           val = obj.dataProcessing(val); // обработка данных перед хагрузкой в таблицу
      });
        $$("bookTable").parse(data.Books);

        data.Staff.forEach(function(val){
          let obj = new modalStaff()
          val = obj.dataProcessing(val);// обработка данных перед хагрузкой в таблицу
          });
        $$("staffTable").parse(data.Staff);
        $$("staffTable").refreshColumns();

        data.Journal.forEach(function(val){
          val.id = val.Id;
          var DateEvent = val.DateEvent.slice(0, 10);
          val.DateEvent = new Date(DateEvent);				
      });
        $$("journalTable").parse(data.Journal);
        $$("journalTable").refreshColumns();
    });
  }

}