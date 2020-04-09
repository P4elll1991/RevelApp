class staffTab{

  constructor(){
    this.modal = new modalStaff();
    this.modal.giveData(this.modal); // загрузка данных в таблицу
  }
    
// кнопки управления
  
  buttons = [{ id:"changeStaff", view:"button", type:"icon", icon:"mdi mdi-account-edit", value: "Изменить",},
  { id:"pushStaff",  view:"button", type:"icon", icon:"mdi mdi-account-plus", value: "Добавить"},
  { id:"goToBook", view:"button", type:"icon", icon:"mdi mdi-book-open-variant", value: "Перейти к книге"},
  { id:"deleteStaff", view:"button", type:"icon", icon:"mdi mdi-account-remove", value: "Удалить"},];
  // колонки таблицы
  columns = [{ id:"ch2", header:{ content:"masterCheckbox", contentId:"mc1" }, template:"{common.checkbox()}", adjust: true,},
  { id:"Name",   header:"ФИО", adjust:true, sort: "int",},
  { id:"Department",  header:"Отдел", adjust: true, sort: "string"},
  { id:"Position",  header:"Должность", adjust:true, sort: "string"},
  { id:"Cellnumber",  header:"Телефон", adjust:true, sort: "int",},
  { id:"BooksStr",  header:"Книги в пользовании", adjust:true, width: 250, sort: "string"},
  ];
    // инициализация таблицы
  init() {
    this.view = {
      view:"layout",
      padding:10,
      id: "staffView", 
      type: "wide",
      
  
      rows: [
        { type: "wide",  
          rows:[ 
  
          { id:"staffSidebar", select:false,// меню
            cols: this.buttons},
          
          // Таблица
  
          {
          view:"datatable", 
          id:"staffTable", 
          wordBreak: "break-all", 
          css:"webix_data_border webix_header_border",
          fixedRowHeight:false,
          multiselect:true, 
          columns: this.columns, 
          data: this.staff, select:true, 
          },
  
        ]},
      ],             
    }
    return this.view;
  }

  initWindow() { // инициализация модального окна
      this.update = new windowStaff();
      this.window = this.update.getWindow();
      return this.window;
    }

  getView() {
    return this.init();
  }
  // прикрепление событий
  editeEvents(parent){

    // событие загрузки таблицы
      $$("staffTable").attachEvent("onresize", function(){
        $$("staffTable").adjustRowHeight(null, true); // формирование высоты таблицы в зависимости от содержимого
      });
      // удаление сотрудника
      $$("deleteStaff").attachEvent("onItemClick", function(){
        parent.delete();
      });
      // изменение сотрудника
      $$("changeStaff").attachEvent("onItemClick", function(){
        parent.checkWin = true;
        var item_data = $$("formStaff").getValues()
        var check = item_data.name;
        if(check != "")
             $$("windowStaff").show();
      });

      // изменение сотрудника двойной щелчок по эелементу

      $$("staffTable").attachEvent("onItemDblClick", function(){
        parent.checkWin = true;
        var item_data = $$("formStaff").getValues()
        var check = item_data.name;
        if(check != "")
             $$("windowStaff").show();
      });

      // добавление сотрудника

      $$("pushStaff").attachEvent("onItemClick", function(){
        parent.checkWin = false;
        $$("formStaff").clear();
        $$("windowStaff").show();
      });

      // событие после выбора элемента

      $$("staffTable").attachEvent("onAfterSelect", function(){
           parent.afterSelect();
      });
// событие после отмены выбора
      $$("staffTable").attachEvent("onAfterUnSelect", function(selection){
          parent.afterUnSelect(selection);
           
      });
//выход из модального окна
      $$("exitWindowStaff").attachEvent("onItemClick", function() {
          $$("windowStaff").hide();
          $$("formStaff").clear();
          $$("formStaff").clearValidation();
          parent.afterSelect();
      });
// обновление данных о сотруднике
      $$("updateStaff").attachEvent("onItemClick", function(){
          parent.updateTab(parent.checkWin, parent);

      });
// нажатие по кнопке к книге
      $$("goToBook").attachEvent("onItemClick", function(){
        parent.focus();
      });
// событие после смены чекбокса
      $$("staffTable").attachEvent("onCheck", function(rowId, colId, state){
        if (state == 1) {
          $$("staffTable").select(rowId, true);
        } else if (state == 0){
          $$("staffTable").unselect(rowId);
        }
      });
  
    }

    // удаление элемента

    delete(){
    var list = $$("staffTable");
    var item_id = list.getSelectedId();
    var item = list.getSelectedItem();
    var IdList = [];
    if (!Array.isArray(item)) {

      // Если один сотрудник
      if (item_id){
        if (item.BooksStr != "") {
          webix.confirm({
            text: "Нельзя удалить сотрудника имеющего задолжность перед библиотекой", 
            ok: "OK",
          }).then(function(){
            return
          });
        } else {
          webix.confirm({
            text: "Вы действительно хотите удалить сотрудника",
            cancel: "Нет", 
            ok: "Да",
          }).then(function(){
            list.remove(item_id);
            IdList.push(item.Id);
            webix.ajax().headers({
              "Content-type":"application/json"
          }).post("/Staff/Delete", JSON.stringify(IdList)).then(function(data){
              data = data.json();
              console.log(data);
              data.forEach(function(val){
                let obj = new modalStaff()
                val = obj.dataProcessing(val);// обработка данных перед загрузкой в таблицу
              });
              $$("staffTable").parse(data);
            });
          });
        }
      } 
    }
    else {  // Если один много сотрудников
      
      var i = 0; 
      item.forEach(function(val){ // перебор массива id
        i++;// счетчик для проверки возможности удаления
        if (val.BooksStr != "") {
          webix.confirm({
            text: "Нельзя удалить сотрудника имеющего задолжность перед библиотекой", 
            ok: "OK",
          }).then(function(){
            return
          });
        } else {
          IdList.push(val.Id);
        }
        
      });
      if (item_id && (IdList.length == i)){ // если счетчик и длина массива совпадают значит все выбранные сотрудники не имеют задолжностей
        webix.confirm({
            text: "Вы действительно хотите удалить сотрудников?",
            cancel: "Нет", 
            ok: "Да",
          }).then(function(){
            list.remove(item_id);
            console.log(IdList);
            webix.ajax().headers({
              "Content-type":"application/json"
          }).post("/Staff/Delete", JSON.stringify(IdList)).then(function(data){
              data = data.json();
              console.log(data);
              data.forEach(function(val){
                let obj = new modalStaff()
                val = obj.dataProcessing(val);// обработка данных перед загрузкой в таблицу
              });
              $$("staffTable").parse(data);
              $$("staffTable").refreshColumns();
            });
          });
      }
    }
  }
//после выбора
  afterSelect() {
      var item = $$("staffTable").getSelectedItem();
      console.log(item);
      $$("formStaff").setValues(item);
      $$("formStaff").setValues(item);
      if (Array.isArray(item)) return;
      item.ch2 = 1; // чекбокс
      
      $$("staffTable").updateItem(item.id, item);
    }

// после отмены выбора
    afterUnSelect(selection){
      var item = selection;
      item.ch2 = 0; // чекбокс
      if(!item.id) return;
      $$("staffTable").updateItem(item.id, item);
    }

// обновление данных о сотруднике

  updateTab(check, parent){
    
     var form = $$("formStaff");
     var item_data = form.getValues();

    form.validate();
    if (!form.validate()){
        webix.message({ type:"error", text:"Некорректно заполненная форма" });
        return
    }
    item_data.Cellnumber = Number(item_data.Cellnumber);
     if(!check) { // если добавляем сотрудника
      parent.addEmployee(item_data)

     } else {
      parent.updateEmp(item_data) // если обновляем сотрудника
     }
     
     
     $$("windowStaff").hide();
     form.clear();
  }
// перевод фокуса к книгам сотрудника
  focus() {
    var item = $$("staffTable").getSelectedItem();
    console.log(item);
    if (!item) return;
    var item_id = item.id;
    var focusId = item.Books;

    if (!focusId) return;
    $$("staffTable").unselect(item_id);
    item.ch2 = 0;
    $$("staffTable").updateItem(item.id, item);
    $$("bookTable").unselectAll();
    

    focusId.forEach(function(v){
        
          $$("bookTable").select(v.Id,true);
          $$("bookTable").showItem(v.Id);
          $$("bookView").show();
      });
    
  }

  // добавить сотрудника

  addEmployee(item_data){
  
      console.log(this.postData);
      webix.ajax().headers({
        "Content-type":"application/json"
    }).post("/Staff/Add", JSON.stringify(item_data)).then(function(data){
      data = data.json();
      console.log(data);
      data.Books.forEach(function(val){
        let obj = new modalBook()
        val = obj.dataProcessing(val);// обработка данных перед загрузкой в таблицу
      });
        $$("bookTable").parse(data.Books);
        data.Staff.forEach(function(val){
          let obj = new modalStaff()
          val = obj.dataProcessing(val);// обработка данных перед загрузкой в таблицу
        });
        let obj = new windowBook()
        obj.optionsBook();
        $$("staffTable").parse(data.Staff);
    });
  }

  // обновить соотрудника

  updateEmp(item_data){
    item_data.Id = Number(item_data.Id);
      
      webix.ajax().headers({
        "Content-type":"application/json"
        }).post("/Staff/Update", JSON.stringify(item_data)).then(function(data){
          data = data.json();
          console.log(data);
          data.Books.forEach(function(val){
            let obj = new modalBook()
            val = obj.dataProcessing(val);
          });
            $$("bookTable").parse(data.Books);
            data.Staff.forEach(function(val){
              let obj = new modalStaff()
              val = obj.dataProcessing(val);
            });
            $$("staffTable").parse(data.Staff);
        });
  }

}