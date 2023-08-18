interface EventPage {
  id: string
  version: number
  name: string
  type: "Event" | "Miscellaneous"
  description: string
  startDate?: Date
  endDate?: Date
  [key: string]: any
}


// TODO: Remove 
// Temporary schema 

