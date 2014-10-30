# Best practices in Android develoment

Lessons learned from Android developers in Futurice. Avoid reinventing the wheel by following these guidelines.

Feedback and criticism are welcomed, feel free to open an issue or send a pull request.

## Summary

#### Use Gradle and its recommended project structure
#### Put passwords and sensitive data in gradle.properties
#### Don't write your own HTTP client, use Volley or Retrofit libraries
#### Use the Jackson library to parse JSON data
#### Use Volley or Retrofit+OkHttp+Picasso for networking and images
#### Avoid Guava and use only a few libraries due to the *65k method limit*
#### Use Fragments to represent a UI screen
#### Use Activities just to manage Fragments
#### Layout XMLs are code, organize them well
#### Use styles to avoid duplicate attributes in layout XMLs
#### Use multiple style files to avoid a single huge one
#### Keep your colors.xml short and DRY, just define the palette
#### Also keep dimens.xml DRY, define generic constants
#### Do not make a deep hierarchy of ViewGroups
#### Avoid client-side processing for WebViews, and beware of leaks
#### Avoid testing with Robolectric on Activities, Fragments, and Views
