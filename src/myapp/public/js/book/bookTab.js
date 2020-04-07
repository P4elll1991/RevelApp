class bookTab {

    constructor(){
    this.modal = new modalBook();
    this.modal.giveData(this.modal);
    
  }
  
    buttons = [
      { id:"change", view:"button", type:"icon", icon:"mdi mdi-pen", value: "Изменить"},                           
      { id:"push",  view:"button", type:"icon", icon:"mdi mdi-plus-box-outline", value: "Добавить"},
      { id:"goToEmployee", view:"button", type:"icon", icon:"mdi mdi-account", value: "Перейти к сотруднику"},
      { id:"delete", view:"button", type:"icon", icon:"mdi mdi-delete-forever", value: "Удалить"},      
      ];
  
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
  
    init() {
    
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

    initWindow() {
      this.up = new windowBook();
      this.window = this.up.getWindow();
      return this.window;
    }

    getView() {
      return this.init();
    }

    editeEvents(parent){
      var options = new windowBook();
      options.optionsBook();
      $$("delete").attachEvent("onItemClick", function(){
        parent.delete();
      });

      $$("change").attachEvent("onItemClick", function(){
        $$("Status").show();
        parent.checkWin = true;
        var item_data = $$("formBook").getValues()
        var check = item_data.BookName;
        console.log(check);
        if(check != "")
             $$("windowBook").show();
      });

      $$("bookTable").attachEvent("onItemDblClick", function(){
        $$("Status").show();
        parent.checkWin = true;
        var item_data = $$("formBook").getValues()
        var check = item_data.name;
        if(check != "")
             $$("windowBook").show();
      });

      $$("push").attachEvent("onItemClick", function(){
        parent.checkWin = false;
        $$("formBook").clear();
        $$("Status").hide();
        $$("formBook").elements["Name"].hide();
        $$("windowBook").show();
      });

      $$("bookTable").attachEvent("onAfterSelect", function(){
           parent.afterSelect();
      });

      $$("bookTable").attachEvent("onAfterUnSelect", function(selection){
          parent.afterUnSelect(selection);
           
      });

      $$("exitWindowBook").attachEvent("onItemClick", function() {
          $$("windowBook").hide();
          $$("formBook").clear();
          $$("formBook").clearValidation();
          $$("Status").show();

          parent.afterSelect();
      });

      $$("updateBookTab").attachEvent("onItemClick", function(){
          
          parent.updateTab(parent.checkWin);

      });

      $$("formBook").elements["Status"].attachEvent("onChange", function(newv, oldv){
        if (newv == "Нет в наличии") {
           $$("formBook").elements["Name"].show();


        } else if (newv == "В наличии"){
          $$("formBook").elements["Name"].hide();
        }
        });


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

    delete(){
    var list = $$("bookTable");
    var item_id = list.getSelectedId();
    var item = list.getSelectedItem();
    console.log(item.Status);
    var IdList = []; 
    
    if (!Array.isArray(item)) {
      if (item_id){
        if (item.Status == "Нет в наличии") {
          webix.confirm({
            text: "Нельзя удалить книгу пока она не будет сдана", 
            ok: "OK",
          }).then(function(){
            return
          });
        } else {
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
                val.id = val.Id;
                  if (val.Employeeid == 1){
                    val.Status = "В наличии";
                    val.Employeei = 0;
                    val.Name = "";
                    val.Cellnumber = null;
                    val.Datestart = null;
                    val.Datefinish = null;
                  } else {
                    val.Status = "Нет в наличии";
                    var Datestart = val.Datestart.slice(0, 10);
                    val.Datestart = new Date(Datestart);
                    val.Datefinish = new Date(Datestart);
                    val.Datefinish.setDate(val.Datefinish.getDate() + 7);
                  }
              });
              $$("bookTable").parse(data);
              });
          });
        } 
      }
    } 
    else {
      var i = 0; 
      item.forEach(function(val){
        i++;
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
      if (item_id && (IdList.length == i)){
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
              val.id = val.Id;
                  if (val.Employeeid == 1){
                    val.Status = "В наличии";
                    val.Employeei = 0;
                    val.Name = "";
                    val.Cellnumber = null;
                    val.Datestart = null;
                    val.Datefinish = null;
                  } else {
                    val.Status = "Нет в наличии";
                    var Datestart = val.Datestart.slice(0, 10);
                    val.Datestart = new Date(Datestart);
                    val.Datefinish = new Date(Datestart);
                    val.Datefinish.setDate(val.Datefinish.getDate() + 7);
                  }
            });
            $$("bookTable").parse(data);
            });
          });
      }
    }
  }

  afterSelect() {
      var item = $$("bookTable").getSelectedItem();
      console.log(item);
      var x = item.Name;
      item.Name = item.Employeeid;
      $$("formBook").setValues(item);
      $$("formBook").setValues(item);
      if (Array.isArray(item)) {
        item.forEach(function(val){
          val.ch1 = 1;
          item.Name = x;
          $$("bookTable").updateItem(val.id, item);
        });
        return;
      }
      item.ch1 = 1;
      item.Name = x;
      $$("bookTable").updateItem(item.id, item);
    }


    afterUnSelect(selection){
      var item = selection;
      item.ch1 = 0;
      if(!item.id) return;
      $$("bookTable").updateItem(item.id, item);
    }

  updateTab(check){
    var table = $$("bookTable");
    var item = table.getSelectedItem();
     var form = $$("formBook");
     var item_data = form.getValues();
     
    

    form.validate();
    if (!form.validate()){
        webix.message({ type:"error", text:"Некорректно заполненная форма" });
        return
    }
     if(!check) {
        if (item_data.id) {
          for (var i in this.books){
          if(item_data.id == this.books[i].id) {
            webix.message({ type:"error", text:"Книга с таким ISBN уже существует" });
            return
          };
        };
      }
      if ((Number(item_data.Year) < 1500 )|| (Number(item_data.Year) > 2100)){
        webix.message({ type:"error", text:"Невалидный год" });
        return
      }
      item_data["Status"] = "В наличии";
      this.postData = {
        action:"info",
        isbn:Number(item_data.Isbn), 
        bookName:item_data.BookName, 
        autor:item_data.Autor, 
        publisher:item_data.Publisher, 
        year:Number(item_data.Year)}
        console.log(this.postData)

      webix.ajax().headers({
        "Content-type":"application/json"
    }).post("/Books/Add", JSON.stringify(this.postData)).then(function(data){
			data = data.json();
			data.forEach(function(val){
				val.id = val.Id;
				if (val.Employeeid == 1){
					val.Status = "В наличии";
					val.Employeei = 0;
					val.Name = "";
					val.Cellnumber = null;
          val.Datestart = null;
          val.Datefinish = null;
				} else {
					val.Status = "Нет в наличии";
					var Datestart = val.Datestart.slice(0, 10);
					val.Datestart = new Date(Datestart);
					val.Datefinish = new Date(Datestart);
					val.Datefinish.setDate(val.Datefinish.getDate() + 7);
				}
			});
      $$("bookTable").parse(data);
		  });


     } else {
      if (item_data.Status) {
        item_data.Employeeid = Number(item_data.Name);
          if (item_data.Status == "В наличии") {
            this.postData = {
              Id:Number(item_data.Id),
              Isbn:Number(item_data.Isbn), 
              BookName:item_data.BookName, 
              Autor:item_data.Autor, 
              Publisher:item_data.Publisher, 
              Year:Number(item_data.Year),
              EmployeeId:1,
      
            }
          item_data.Name = "";
          item_data.Datestart = "";
          item_data.Datefinish = "";
        } else {
          this.postData = {
            Id:Number(item_data.Id),
            Isbn:Number(item_data.Isbn), 
            BookName:item_data.BookName, 
            Autor:item_data.Autor, 
            Publisher:item_data.Publisher, 
            Year:Number(item_data.Year),
            EmployeeId:item_data.Employeeid,
    
          }
          var today = new Date;
          var dateFinish = new Date;
          item_data.Datestart = new Date;
          dateFinish.setDate(dateFinish.getDate() + 7);
          item_data.Datefinish = dateFinish;

        }
      }

        $$("staffTable").eachRow(function(row){
          var record = $$("staffTable").getItem(row);
          if (record.Id == item_data.Name){
            item_data.Cellnumber = record.Cellnumber;
            item_data.Name = record.Name;
          }
      });
      
        console.log(item_data);
        if (item.Status != item_data.Status){
          console.log("нет");
          if(item_data.Status == "В наличии"){
            console.log("Возвращено")
            this.postDataEvent = {
              Event: "Возвращено",
              BookId :Number(item_data.Id),
              BookName: item_data.BookName,
              Isbn: Number(item_data.Isbn),
              EmployeeId: Number(item.Employeeid),
              Name: item.Name,
              Cellnumber: Number(item.Cellnumber)
            };
            console.log(this.postDataEvent)
          } else {
            console.log("Выдано")
            this.postDataEvent = {
              Event: "Выдано",
              BookId :Number(item_data.Id),
              BookName: item_data.BookName,
              Isbn: Number(item_data.Isbn),
              Name: item_data.Name,
              EmployeeId: Number(item_data.Employeeid),
              Cellnumber: Number(item_data.Cellnumber),
            };
            console.log(this.postDataEvent)
          }

          webix.ajax().headers({
            "Content-type":"application/json"
        }).post("/Journal/Add", JSON.stringify(this.postDataEvent));



        } 
        
          console.log(this.postData);
          webix.ajax().headers({
            "Content-type":"application/json"
        }).post("/Books/Update", JSON.stringify(this.postData)).then(function(data){
          data = data.json();
          console.log(data);
          data.Books.forEach(function(val){
                val.id = val.Id;
            if (val.Employeeid == 1){
              val.Status = "В наличии";
              val.Employeei = 0;
              val.Name = "";
              val.Cellnumber = null;
              val.Datestart = null;
              val.Datefinish = null;
            } else {
              val.Status = "Нет в наличии";
              var Datestart = val.Datestart.slice(0, 10);
              val.Datestart = new Date(Datestart);
              val.Datefinish = new Date(Datestart);
              val.Datefinish.setDate(val.Datefinish.getDate() + 7);
				}
          });
            $$("bookTable").parse(data.Books);
            data.Staff.forEach(function(val){
              val.id = val.Id;
              val.BooksStr = "";
              val.Books.forEach(function(v){
                val.BooksStr += "<p style = 'padding: 0px; margin: 0px; height: 25px;'> ISBN : " + v.Isbn + ", " + v.BookName + ", " + v.Datestart + " - " + v.Datefinish + ";</p>";
              });
            });
            $$("staffTable").parse(data.Staff);
            $$("journalTable").parse(data.Journal);
        });
     }
     
     
     $$("windowBook").hide();
     form.clear();
  }

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


}