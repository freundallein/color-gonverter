# color-gonverter
Golang gRPC service for converting color-points from ARGB to RGBA (32/64bit) 

Have 2 RPC calls ConvertARGB_RGBA32 and ConvertARGB_RGBA64,  
Work via gRPC stream.  
Typical message:  
`message ARGB32 {
    uint32 argb = 1;
}`
Response:
`message RGBA32 {
    uint32 rgba = 1;
}`
