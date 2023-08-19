interface EventPage {
  id: string
  version: number
  name: string
  type: "Event" | unknown
  description: string
  startDate?: Date
  endDate?: Date
  createdOn?: Date
  lastModified?: Date
  [key: string]: any
}


// TODO: Remove 
// Temporary schema 

