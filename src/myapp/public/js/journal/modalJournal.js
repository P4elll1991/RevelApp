class modalJournal {

	

	takeData(){

		webix.ajax().get("/Journal/Give").then(function(data){
			data = data.json();
			data.forEach(function(val){
				val.id = val.Id;
				var DateEvent = val.DateEvent.slice(0, 10);
				val.DateEvent = new Date(DateEvent);				
			});
			console.log(data);
            $$("journalTable").parse(data);
		  });

	}

	giveData(parent) {
		parent.takeData();
	}
}