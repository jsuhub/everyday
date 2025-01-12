char *restoreString(char *s, int *indices, int indicesSize)
{
    char* newArr = (char*)malloc(sizeof(char) * (indicesSize + 1));
    for(int i = 0; i < indicesSize; i++) {
       newArr[indices[i]] = s[i];
    }
    newArr[indicesSize] = '\0';
    return newArr;
}
