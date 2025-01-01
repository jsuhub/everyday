
#include <stdio.h>
#include <string.h>
#include <assert.h>
int lengthOfLastWord(char *s)
{
    int lastWordLastChar = 0;
    for (int i = strlen(s) - 1; i >= 0; i--)
    {
        if (s[i] != ' ')
        {
            lastWordLastChar = i;
            printf("%d", lastWordLastChar);
            break;
        }
    }
    int lastWordFirstChar = 0;
    for (int i = lastWordLastChar; i >= 0; i--)
    {
        if (s[i] == ' ')
        {
            lastWordFirstChar = i + 1;
            printf("%d", lastWordFirstChar);
            break;
        }
    }
    return lastWordLastChar - lastWordFirstChar + 1;
}
int main()
{
    assert(lengthOfLastWord(" hello demo") == 4);
}
