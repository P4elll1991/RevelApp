class journalTab {


    constructor(){
    this.modal = new modalJournal();
    this.modal.giveData(this.modal);
  }
  // кнопки управления
    buttons = [
      { id:"goToEmployeeLog", view:"button", type:"icon", icon:"mdi mdi-account", value: "Перейти к сотруднику"},
      { id:"goToBookLog", view:"button", type:"icon", icon:"mdi mdi-book-open-variant", value: "Перейти к книге"},
    ];
  
    columns = [
      { id:"Event",    header:"Событие",  sort: "string",  adjust:true,},
      { id:"BookNameJ",   header:"Название",   sort: "string",  adjust:true,},
      { id:"IsbnJ",  header:"ISBN",   sort: "int",  adjust:true,},
      { id:"DateEvent",  header:"Дата события",  format:webix.Date.dateToStr("%d.%m.%Y"), sort: "date",  adjust:true,},
      { id:"NameJ",  header:"ФИО",   sort: "string",  adjust:true,},
      { id:"CellnumberJ",  header:"Телефон",   sort: "string",  adjust:true,},
      ];
  // инициализация таблицы
    init(){
      this.view = {
        view:"layout",
        padding:10,
        id: "journalView", 
        type: "wide",
    
        rows: [
          { type: "wide",  
            rows:[ 
    
            { id:"journalSidebar", select:false,// меню
              cols: this.buttons},
            
            // Таблица
    
            {
            view:"datatable", 
            id:"journalTable", 
            wordBreak: "break-all", 
            css:"webix_data_border webix_header_border",
            multiselect:true, 
            columns: this.columns, 
            select:true, 
            },
    
          ]},
        ],             
      }
      return this.view;
    }

    getView() {
      return this.init();
    }
// добавление событий
    editeEvents(parent){

      $$("goToBookLog").attachEvent("onItemClick", function(){
        parent.focusBook(); // фокус к ниге
      });

      $$("goToEmployeeLog").attachEvent("onItemClick", function(){
        parent.focusStaff("staffTable"); // фокус к сотруднику
      });
  
    }
// функция фокуса к книге
    focusBook() {
      var item = $$("journalTable").getSelectedItem();
      
      if (!item) return;
      var item_id = item.id;
      var focusId = item.BookId;
      var flagCheck = false;
      $$("bookTable").eachRow(function(row){
        var record = $$("bookTable").getItem(row)
        if (focusId == record.Id) {
          flagCheck = true
        }
      });
      if (!flagCheck) {
        webix.message({ type:"error", text:"Книга была удален" });
        return
      };

      if (!focusId) return;

      $$("journalTable").unselect(item_id);
      $$("bookTable").unselectAll();
      $$("bookTable").select(focusId,true);
      $$("bookView").show();
      $$("bookTable").showItem(focusId);

    }

// функция фокуса к книге
    focusStaff() {
      var item = $$("journalTable").getSelectedItem();
      if (!item) return;
      var item_id = item.id;
      var focusId = item.EmployeeId;
      var flagCheck = false;
      $$("staffTable").eachRow(function(row){
        var record = $$("staffTable").getItem(row)
        if (focusId == record.Id) {
          flagCheck = true
        }
      });
      if (!flagCheck) {
        webix.message({ type:"error", text:"Сотрудник был удален" });
        return
      };

      $$("journalTable").unselect(item_id);
      $$("staffTable").unselectAll();
      $$("staffTable").select(focusId,true);
      $$("staffView").show();
      $$("staffTable").showItem(focusId);
    }
    
  }
    