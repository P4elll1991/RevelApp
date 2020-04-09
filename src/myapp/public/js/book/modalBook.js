class modalBook {

	

	takeData(parent){
		
		webix.ajax().get("/Books/Give").then(function(data){
			data = data.json();
			console.log(data);
			data.forEach(function(val){
				val = parent.dataProcessing(val);
			});
            $$("bookTable").parse(data);
		  });
		
	}

	giveData(parent) {
		parent.takeData(parent);
	}

	//функция обработки данных перед загрузкой на страницу

	dataProcessing(val) {
		val.id = val.Id;
		if (val.Employeeid == 1){
			val.Status = "В наличии";
			val.Employeei = 0;
			val.Name = "";
			val.Cellnumber = null;
			val.Datestart = null;
		} else {
			val.Status = "Нет в наличии";
			var Datestart = val.Datestart.slice(0, 10);
			val.Datestart = new Date(Datestart);
			val.Datefinish = new Date(Datestart);
			val.Datefinish.setDate(val.Datefinish.getDate() + 7);
		}
		return val;
	}

}