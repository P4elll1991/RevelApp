class modalStaff {

	

	takeData(){
		
		webix.ajax().get("/Staff/Give").then(function(data){
			data = data.json();
			
			data.forEach(function(val){
				val.id = val.Id;
				val.BooksStr = "";
				val.Books.forEach(function(v){
					var Datestart = v.Datestart.slice(0, 10);
					v.Datestart = new Date(Datestart);
					v.Datefinish = new Date(Datestart);
					v.Datefinish.setDate(v.Datefinish.getDate() + 7);
					var format = webix.Date.dateToStr("%d.%m.%Y");
					val.BooksStr += "<p style = 'padding: 0px; margin: 0px; height: 25px;'> ISBN : " + v.Isbn + ", " + v.BookName + ", " + format(v.Datestart) + " - " + format(v.Datefinish) + ";</p>";
				});
			});
			$$("staffTable").parse(data);
			
		  });
	}

	giveData(parent) {
		parent.takeData();
	}
}